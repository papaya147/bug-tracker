package main

import (
	"log"

	"github.com/papaya147/bug-tracker/backend/api"
	"github.com/papaya147/bug-tracker/backend/config"
	db "github.com/papaya147/bug-tracker/backend/db/sqlc"
	"github.com/papaya147/bug-tracker/backend/util"
)

func main() {
	config := config.NewConfig()

	postgresConn := util.CreatePostgresPool(config.POSTGRES_DSN)

	store := db.NewStore(postgresConn)

	server := api.NewServer(store)

	log.Println("starting http server on port", config.HTTP_SERVER_PORT)
	server.Start(config.HTTP_SERVER_PORT)
}
