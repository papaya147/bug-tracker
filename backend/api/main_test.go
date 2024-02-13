package api

import (
	"log"
	"os"
	"testing"

	"github.com/papaya147/buggy/backend/api/profile"
)

var testApp *server

func TestMain(m *testing.M) {
	testApp := &server{}

	testApp.profileHandler = profile.NewHandler(nil, nil, nil)
	testApp.router = testApp.routes()
	log.Println(testApp.router)

	os.Exit(m.Run())
}
