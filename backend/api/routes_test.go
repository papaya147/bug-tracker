package api

import (
	"testing"

	"github.com/papaya147/buggy/backend/test"
)

func TestRoutes(t *testing.T) {
	router := testApp.router
	apiVersion := 1

	apiRoutes := []test.Route{
		{Path: "/profile/", Method: "POST"},
		{Path: "/profile/verify", Method: "GET"},
		{Path: "/profile/login", Method: "POST"},
		{Path: "/profile/", Method: "GET"},

		{Path: "/organisation/", Method: "POST"},
		{Path: "/organisation/", Method: "GET"},
		{Path: "/organisation/", Method: "PUT"},
		{Path: "/organisation/transfer/", Method: "POST"},
		{Path: "/organisation/transfer/", Method: "GET"},
		{Path: "/organisation/transfer/{organisation-transfer-id}", Method: "DELETE"},
		{Path: "/organisation/transfer/response/{organisation-transfer-id}", Method: "GET"},
		{Path: "/organisation/team/", Method: "POST"},
		{Path: "/organisation/team/", Method: "GET"},
		{Path: "/organisation/team/{team-id}", Method: "PUT"},

		{Path: "/team-member/", Method: "POST"},
		{Path: "/team-member/{team-id}", Method: "GET"},
		{Path: "/team-member/", Method: "PUT"},
	}

	for _, apiRoute := range apiRoutes {
		test.RouteExists(t, router, apiVersion, apiRoute)
	}
}
