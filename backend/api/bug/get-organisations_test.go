package bug

import (
	"net/http"
	"testing"

	"github.com/papaya147/buggy/backend/test"
)

func TestGetOrganisations(t *testing.T) {
	test.TestCase(t, http.MethodGet, "/organisations", testHandler.getOrganisations, nil, http.StatusOK)
}
