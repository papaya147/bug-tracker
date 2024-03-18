package bug

import (
	"encoding/json"
	"net/http"
	"testing"

	db "github.com/papaya147/buggy/backend/db/sqlc"
	"github.com/papaya147/buggy/backend/test"
	"github.com/papaya147/buggy/backend/util"
)

func TestUpdate(t *testing.T) {
	postBody := map[string]interface{}{
		"id":          util.RandomUuid(),
		"name":        util.RandomString(10),
		"description": util.RandomString(20),
		"status":      db.BugstatusPENDING,
		"priority":    db.BugpriorityHIGH,
	}
	body, _ := json.Marshal(postBody)
	test.TestCase(t, http.MethodPut, "/", testHandler.update, body, http.StatusOK)
}
