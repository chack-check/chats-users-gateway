package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/chack-check/chats-users-gateway/infrastructure/api/graph"
	"github.com/chack-check/chats-users-gateway/infrastructure/api/middlewares"
	"github.com/chack-check/chats-users-gateway/infrastructure/api/settings"
	"github.com/go-chi/chi"
)

func RunApi() {
	router := chi.NewRouter()

	router.Use(middlewares.UserMiddleware)
	router.Use(middlewares.CorsMiddleware)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/api/v1/chats-users", playground.Handler("GraphQL playground", "/api/v1/chats-users/query"))
	router.Handle("/api/v1/chats-users/query", srv)

	listen := fmt.Sprintf(":%d", settings.Settings.APP_PORT)
	log.Printf("Running server on %d port", settings.Settings.APP_PORT)
	log.Fatal(http.ListenAndServe(listen, router))
}
