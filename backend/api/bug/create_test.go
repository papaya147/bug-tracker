package bug

import (
	"encoding/json"
	"net/http"
	"testing"

	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/test"
	"github.com/papaya147/buggy/backend/util"
)

func TestCreate(t *testing.T) {
	postBody := map[string]interface{}{
		"name":          util.RandomString(10),
		"description":   util.RandomString(10),
		"priority":      db.BugpriorityHIGH,
		"assigned_team": util.RandomUuid(),
		"assignee_team": util.RandomUuid(),
	}
	body, _ := json.Marshal(postBody)
	test.TestCase(t, http.MethodPost, "/", testHandler.create, body, http.StatusOK)
}
