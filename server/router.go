package server

import (
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/M1RCLE/comments/graph"
	"github.com/M1RCLE/comments/graph/generated"
	contracts "github.com/M1RCLE/comments/src/service/interface"
	"github.com/rs/zerolog/log"
)

type Router struct {
	Mux *http.ServeMux
}

func NewRouter(commentService contracts.Comment, postService contracts.Post, subscriptionService contracts.Subscription) *Router {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: graph.NewResolver(postService, commentService, subscriptionService),
	}))

	m := http.NewServeMux()

	m.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	m.Handle("/graphql", logRequest(srv))

	return &Router{
		Mux: m,
	}
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Info().Msgf("Started %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Info().Msgf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}
