package profile

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/papaya147/buggy/backend/test"
)

func TestChangePassword(t *testing.T) {
	postBody := map[string]interface{}{
		"old_password": "some password",
		"new_password": "some new password",
	}
	body, _ := json.Marshal(postBody)
	test.TestCase(t, "POST", "/password/change", testHandler.changePassword, body, http.StatusOK)
}
