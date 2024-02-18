package transfer

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/papaya147/buggy/backend/test"
)

func TestCreate(t *testing.T) {
	postBody := map[string]interface{}{
		"email": "test@test.com",
	}
	body, _ := json.Marshal(postBody)
	test.TestCase(t, "POST", "/", testHandler.create, body, http.StatusOK)
}
