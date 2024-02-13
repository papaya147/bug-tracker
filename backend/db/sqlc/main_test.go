package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/papaya147/buggy/backend/config"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config := config.NewConfig("../../")
	log.Println(config.POSTGRES_DSN)

	conn, err := pgxpool.New(context.Background(), config.POSTGRES_DSN)
	if err != nil {
		log.Panicf("unable to connect to the database... %s", err.Error())
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
