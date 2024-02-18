package transfer

import (
	"net/http"
	"testing"

	"github.com/papaya147/buggy/backend/test"
)

func TestDelete(t *testing.T) {
	test.TestCase(t, "DELETE", "/", testHandler.delete, nil, http.StatusOK)
}
