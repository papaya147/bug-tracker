package team

import (
	"os"
	"testing"

	"github.com/papaya147/buggy/backend/config"
	db "github.com/papaya147/buggy/backend/db/sqlc"
)

var testHandler *Handler

func TestMain(m *testing.M) {
	config := config.NewConfig("../../../")

	store := db.NewMockStore()

	testHandler = NewHandler(config, store)

	os.Exit(m.Run())
}
