package organisation

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/papaya147/buggy/backend/test"
)

func TestUpdate(t *testing.T) {
	postBody := map[string]interface{}{
		"name":        "Test",
		"description": "Test",
	}
	body, _ := json.Marshal(postBody)
	test.TestCase(t, "POST", "/", testHandler.update, body, http.StatusOK)
}
