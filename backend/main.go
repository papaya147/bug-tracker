package main

import (
	"context"
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/papaya147/buggy/backend/api"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

// @title Buggy APIs
// @version 1.0
// @description The comprehensive list of all Buggy APIs
// @host localhost:4000
// @BasePath /api/v1
func main() {
	config := util.NewConfig("./")

	postgresConn := util.CreatePostgresPool(config.POSTGRES_DSN)
	defer postgresConn.Close()

	util.CreateCookieStore(config)

	log.Print("attempting database migration...")
	if err := runDbMigration(config); err != nil {
		log.Fatal("database migration failed with error:", err)
	} else {
		log.Println("database migration was successful!")
	}

	store := db.NewStore(postgresConn)

	tokenMaker, err := token.NewPasetoMaker(context.Background(), "0123456789012345678901234567890123456789012345678901234567890123")
	if err != nil {
		log.Fatal("cannot create token maker:", err)
	}

	server := api.NewServer(store, tokenMaker)

	log.Println("starting http server on port", config.HTTP_SERVER_PORT)
	server.Start(config.HTTP_SERVER_PORT)
}

func runDbMigration(config util.Config) error {
	migration, err := migrate.New(config.MIGRATION_URL, config.POSTGRES_DSN)
	if err != nil {
		return err
	}

	if err := migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
