package util

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePostgresPool(dsn string) *pgxpool.Pool {
	count := 0
	for {
		db, err := pgxpool.New(context.Background(), dsn)
		if err != nil {
			count++
		} else {
			log.Println("connected to postgres!")
			return db
		}
		if count == 5 {
			log.Println("unable to connect to postgres...", err)
			log.Println("retying in 5 seconds...")
			time.Sleep(time.Second * 5)
			count = 0
		}
	}
}
