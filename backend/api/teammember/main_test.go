package teammember

import (
	"os"
	"testing"

	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

var testHandler *Handler

func TestMain(m *testing.M) {
	config := util.NewConfig("../../")

	store := db.NewMockStore()

	maker := token.NewMockMaker()

	testHandler = NewHandler(config, store, maker)

	os.Exit(m.Run())
}
