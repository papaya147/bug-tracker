package api

import (
	"testing"

	"github.com/papaya147/buggy/backend/test"
)

func TestRoutes(t *testing.T) {
	testRoutes := testApp.router
	apiVersion := 1

	apiRoutes := []test.Route{
		{Path: "/profile/", Method: "POST"},
		{Path: "/profile/verify", Method: "GET"},
		{Path: "/profile/login", Method: "POST"},
		{Path: "/profile/", Method: "GET"},
	}

	for _, apiRoute := range apiRoutes {
		test.RouteExists(t, testRoutes, apiVersion, apiRoute)
	}
}
