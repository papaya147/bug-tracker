package api

import (
	"os"
	"testing"

	"github.com/papaya147/buggy/backend/api/organisation"
	"github.com/papaya147/buggy/backend/api/profile"
	"github.com/papaya147/buggy/backend/api/teammember"
	"github.com/papaya147/buggy/backend/util"
)

var testApp *server

func TestMain(m *testing.M) {
	testApp = &server{}

	testApp.profileHandler = profile.NewHandler(util.Config{}, nil, nil)
	testApp.organisationHandler = organisation.NewHandler(util.Config{}, nil, nil)
	testApp.teamMemberHandler = teammember.NewHandler(util.Config{}, nil, nil)
	testApp.router = testApp.routes()

	os.Exit(m.Run())
}
