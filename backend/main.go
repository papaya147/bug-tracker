package main

import (
	"context"
	"log"

	"github.com/papaya147/buggy/backend/api"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

func main() {
	config := util.NewConfig(".")

	postgresConn := util.CreatePostgresPool(config.POSTGRES_DSN)

	store := db.NewStore(postgresConn)

	tokenMaker, err := token.NewPasetoMaker(context.Background(), "0123456789012345678901234567890123456789012345678901234567890123")
	if err != nil {
		log.Fatal("cannot create token maker:", err)
	}

	server := api.NewServer(store, tokenMaker)

	log.Println("starting http server on port", config.HTTP_SERVER_PORT)
	server.Start(config.HTTP_SERVER_PORT)
}
