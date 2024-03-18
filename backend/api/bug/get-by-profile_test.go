package bug

import (
	"net/http"
	"testing"

	"github.com/papaya147/buggy/backend/test"
)

func TestGetByProfile(t *testing.T) {
	test.TestCase(t, http.MethodGet, "/by-profile", testHandler.getByProfile, nil, http.StatusOK)
}
