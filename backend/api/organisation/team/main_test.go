package team

import (
	"os"
	"testing"

	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/util"
)

var testHandler *Handler

func TestMain(m *testing.M) {
	config := util.NewConfig("../../../")

	store := db.NewMockStore()

	testHandler = NewHandler(config, store)

	os.Exit(m.Run())
}
