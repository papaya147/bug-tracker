package transfer

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/papaya147/buggy/backend/test"
	"github.com/papaya147/buggy/backend/util"
)

func TestResponse(t *testing.T) {
	test.TestCase(t, "GET", fmt.Sprintf("/response/%s?status=%t", util.RandomUuid(), util.RandomBool()), testHandler.response, nil, http.StatusOK)
}
