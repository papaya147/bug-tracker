package organisation

import (
	"os"
	"testing"

	"github.com/papaya147/buggy/backend/config"
	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
)

var testHandler *Handler

func TestMain(m *testing.M) {
	config := config.NewConfig("../../")

	store := db.NewMockStore()

	maker := token.NewMockMaker()

	testHandler = NewHandler(config, store, maker)

	os.Exit(m.Run())
}
