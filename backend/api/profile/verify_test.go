package profile

import (
	"net/http"
	"testing"

	"github.com/papaya147/buggy/backend/test"
)

func TestVerify(t *testing.T) {
	test.TestCase(t, "GET", "/verify", testHandler.verify, nil, http.StatusOK)
}
