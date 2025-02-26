package test

import (
	"testing"

	"github.com/M1RCLE/comments/server"
	inmemory "github.com/M1RCLE/comments/src/repository/inmemory"
	"github.com/M1RCLE/comments/src/service"
	"github.com/rs/zerolog/log"
)

func TestInMemoryPostCreation(t *testing.T) {
	repository := inmemory.NewStorage()
	log.Info().Msg("Repository started...")

	commentService := service.NewCommentService(repository)
	postService := service.NewPostService(repository)
	subscriptionService := service.NewSubscriptionService(repository)

	router := server.NewRouter(commentService, postService, subscriptionService)

	controller := server.NewMediaController(router.Mux, "8080")
	controller.LinearStart()
}
