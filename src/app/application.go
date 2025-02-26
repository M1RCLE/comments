package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	"github.com/M1RCLE/comments/server"
	"github.com/M1RCLE/comments/src/repository/contract"
	inmemory "github.com/M1RCLE/comments/src/repository/inmemory"

	"github.com/M1RCLE/comments/src/config"
	"github.com/M1RCLE/comments/src/service"
)

func StartApp() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config file")
	}

	log.Logger = setupLog()

	var repository contract.Repository

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errGroup, errCtx := errgroup.WithContext(ctx)

	switch cfg.RepositoryType {
	case "postgres":
		// commentStorage, err := db.NewStorage[entity.Comment](*cfg)
		// if err != nil {
		// 	log.Fatal().Err(err).Msg("Failed to create comment storage")
		// }

		// log.Info().Msg("Comment storage connected...")

		// errGroup.Go(func() error {
		// 	<-errCtx.Done()
		// 	if err = commentStorage.Close(); err != nil {
		// 		return fmt.Errorf("failed to close comment storage %w", err)
		// 	}

		// 	log.Info().Msg("Comment storage disconnected...")
		// 	return nil
		// })

		// postStorage, err := db.NewStorage[entity.Post](*cfg)
		// if err != nil {
		// 	log.Fatal().Err(err).Msg("Failed to create post storage")
		// }

		// log.Info().Msg("Post storage connected...")

		// errGroup.Go(func() error {
		// 	<-errCtx.Done()
		// 	if err = postStorage.Close(); err != nil {
		// 		return fmt.Errorf("failed to close comment storage %w", err)
		// 	}

		// 	log.Info().Msg("Post storage disconnected...")
		// 	return nil
		// })

		// MediaRepository = db.NewDatabaseMedia(commentStorage, postStorage)

	case "in_memory":
		log.Info().Msg("In-memory storage connected...")
		repository = inmemory.NewStorage()
	default:
		log.Fatal().Msg("Unknown repository type")
	}

	log.Info().Msg("Repository started...")

	commentService := service.NewCommentService(repository)
	postService := service.NewPostService(repository)
	subscriptionService := service.NewSubscriptionService(repository)

	router := server.NewRouter(commentService, postService, subscriptionService)

	controller := server.NewMediaController(router.Mux, cfg.Port)
	controller.Start()
	log.Info().Msg("Controller started...")

	log.Info().Msg("App started...")

	errGroup.Go(func() error {
		<-errCtx.Done()

		if err = controller.Shutdown(); err != nil {
			return fmt.Errorf("failed to shutdown controller: %w", err)
		}
		log.Info().Msg("Media controller stopped...")
		return nil
	})

	errGroup.Go(func() error {
		<-errCtx.Done()

		timeoutCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()

		select {
		case <-timeoutCtx.Done():
			if errors.Is(timeoutCtx.Err(), context.DeadlineExceeded) {
				return timeoutCtx.Err()
			}

			return nil
		}
	})

	err = errGroup.Wait()

	if errors.Is(err, context.Canceled) || err == nil {
		log.Info().Msg("App stopped...")
	} else if err != nil {
		log.Fatal().Err(err).Msg("Failed to shutdown app")
	}
}

func setupLog() zerolog.Logger {
	return zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
}
