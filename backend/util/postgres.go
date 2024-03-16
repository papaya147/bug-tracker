package util

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
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

func CreateDatabase(conn *pgxpool.Pool) {
	count := 0
	for {
		query := `CREATE DATABASE "buggy";`
		_, err := conn.Exec(context.Background(), query)
		if err != nil {
			if e, ok := err.(*pgconn.PgError); ok && e.Code == "42P04" {
				log.Println("database already exists...")
				return
			}
			count++
		} else {
			log.Println("database created!")
			return
		}
		if count == 5 {
			log.Println("unable to create database...", err)
			log.Println("retying in 5 seconds...")
			time.Sleep(time.Second * 5)
			count = 0
		}
	}
}
