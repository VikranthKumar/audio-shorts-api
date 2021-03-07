package main

import (
	"context"
	"github.com/golang-migrate/migrate/v4"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nooble/task/audio-short-api/pkg/api"
	"github.com/nooble/task/audio-short-api/pkg/api/generated"
	"github.com/nooble/task/audio-short-api/pkg/config"
	"github.com/nooble/task/audio-short-api/pkg/db"
	"github.com/nooble/task/audio-short-api/pkg/logging"
	"github.com/nooble/task/audio-short-api/pkg/store"
	"github.com/nooble/task/audio-short-api/pkg/util"
)

func main() {
	ctx := context.Background()

	// =========== logger ============= //
	ctx = logging.NewContext(ctx)

	// =========== config ============= //
	cfg, err := config.New()
	util.ExitOnErr(ctx, err)

	// =========== db ============= //
	pgDB, err := db.NewDB(cfg)
	util.ExitOnErr(ctx, err)
	defer pgDB.Close()

	// =========== migrator ============== //
	m, err := db.NewMigrator(pgDB)
	util.ExitOnErr(ctx, err)
	defer m.Close()

	// up migrations are done here; exit on failed migrations
	go func() {
		err := m.Up()
		if err != nil && err != migrate.ErrNoChange { // ignore ErrNoChange, meaning it is at the latest version
			m.GracefulStop <- true
			util.ExitOnErr(ctx, err)
		}
	}()

	// =========== datastore ============= //
	asStore, err := store.NewShortsStore(pgDB)
	util.ExitOnErr(ctx, err)

	cStore, err := store.NewCreatorsStore(pgDB)
	util.ExitOnErr(ctx, err)

	// =========== resolver ============= //
	resolver, err := api.New(asStore, cStore)
	util.ExitOnErr(ctx, err)

	// =========== server ============= //
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logging.WithContext(ctx).Info("connected for GraphQL playground")
	err = http.ListenAndServe(cfg.Server.Host+":"+cfg.Server.Port, nil)
	util.ExitOnErr(ctx, err)
}
