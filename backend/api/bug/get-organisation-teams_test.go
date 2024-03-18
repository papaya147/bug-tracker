package bug

import (
	"net/http"
	"testing"

	"github.com/papaya147/buggy/backend/test"
)

func TestGetOrganisationTeams(t *testing.T) {
	test.TestCase(t, http.MethodGet, "/", testHandler.getOrganisationTeams, nil, http.StatusForbidden)
}
