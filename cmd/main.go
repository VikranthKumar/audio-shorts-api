package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nooble/task/audio-short-api/pkg/api"
	"github.com/nooble/task/audio-short-api/pkg/api/generated"
	"net/http"
	"os"

	"github.com/nooble/task/audio-short-api/pkg/db"
	"github.com/nooble/task/audio-short-api/pkg/logging"
	"github.com/nooble/task/audio-short-api/pkg/store"
	"github.com/nooble/task/audio-short-api/pkg/util"
)

const defaultPort = "8080"

func main() {
	ctx := context.Background()

	// =========== config ============= //
	//TODO add config

	// =========== logger ============= //
	ctx = logging.NewContext(ctx)

	// =========== db ============= //
	pgDB, err := db.New("")
	util.ExitOnErr(ctx, err)
	defer pgDB.Close()

	// =========== datastore ============= //
	asStore, err := store.New(pgDB)
	util.ExitOnErr(ctx, err)

	// =========== resolver ============= //
	resolver, err := api.New(asStore)
	util.ExitOnErr(ctx, err)

	// =========== server ============= //
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logging.WithContext(ctx).Info("connected for GraphQL playground")
	err = http.ListenAndServe(":"+port, nil)
	util.ExitOnErr(ctx, err)
}
