package profile

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/papaya147/buggy/backend/test"
)

func TestCreate(t *testing.T) {
	postBody := map[string]interface{}{
		"name":     "papaya",
		"email":    "test@test.com",
		"password": "some password",
	}
	body, _ := json.Marshal(postBody)
	test.TestCase(t, "POST", "/", testHandler.create, body, http.StatusOK)
}
