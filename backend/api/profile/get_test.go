package profile

import (
	"net/http"
	"testing"

	"github.com/papaya147/buggy/backend/test"
)

func TestGet(t *testing.T) {
	test.TestCase(t, "GET", "/", testHandler.get, nil, http.StatusOK)
}
