package teammember

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/papaya147/buggy/backend/test"
	"github.com/papaya147/buggy/backend/util"
)

func TestUpdate(t *testing.T) {
	postBody := map[string]interface{}{
		"admin":      util.RandomBool(),
		"team_id":    util.RandomUuid(),
		"profile_id": util.RandomUuid(),
	}
	body, _ := json.Marshal(postBody)
	test.TestCase(t, "PUT", "/", testHandler.update, body, http.StatusOK)
}
