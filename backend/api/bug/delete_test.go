package bug

import (
	"net/http"
	"testing"

	"github.com/papaya147/buggy/backend/test"
)

func TestDelete(t *testing.T) {
	test.TestCase(t, http.MethodDelete, "/", testHandler.delete, nil, http.StatusOK)
}
