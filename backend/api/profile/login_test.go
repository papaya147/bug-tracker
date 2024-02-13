package profile

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/papaya147/buggy/backend/test"
)

func TestLogin(t *testing.T) {
	postBody := map[string]interface{}{
		"email":    "test@test.com",
		"password": "some password",
	}
	body, _ := json.Marshal(postBody)
	test.TestCase(t, "POST", "/login", testHandler.login, body, http.StatusOK)
}
