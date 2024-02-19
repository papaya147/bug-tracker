package teammember

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/papaya147/buggy/backend/test"
)

func TestCreate(t *testing.T) {
	postBody := map[string]interface{}{
		"team_id": uuid.New(),
		"email":   "test@test.com",
		"admin":   false,
	}
	body, _ := json.Marshal(postBody)
	test.TestCase(t, "POST", "/", testHandler.create, body, http.StatusOK)
}
