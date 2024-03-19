package bug

import (
	"net/http"
	"testing"

	"github.com/papaya147/buggy/backend/test"
)

func TestGetOrganisationTeamsByProfile(t *testing.T) {
	test.TestCase(t, http.MethodGet, "/", testHandler.getOrganisationTeamsByProfile, nil, http.StatusOK)
}
