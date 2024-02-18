package test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/papaya147/buggy/backend/token"
	"github.com/papaya147/buggy/backend/util"
)

func TestCase(t *testing.T, method, route string, handler http.HandlerFunc, body []byte, expected int, headers ...string) {
	req, _ := http.NewRequest(method, route, bytes.NewReader(body))
	rr := httptest.NewRecorder()

	req = addChiURLParams(req, map[string]string{
		"organisation-transfer-id": util.RandomUuid().String(),
	})

	for _, header := range headers {
		parts := strings.Split(header, ":")
		if len(parts) != 2 {
			continue
		}
		req.Header.Add(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
	}

	req = req.WithContext(addMockTokenToContext(req.Context()))
	handler.ServeHTTP(rr, req)

	if rr.Code != expected {
		t.Errorf("expected status code %d, got %d with body %s", expected, rr.Code, rr.Body)
	}
}

func addMockTokenToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, token.TokenPayloadKey, token.Payload{
		UserId:    util.RandomUuid(),
		TokenId:   util.RandomUuid(),
		TokenType: token.AccessToken,
	})
}

func addChiURLParams(r *http.Request, params map[string]string) *http.Request {
	ctx := chi.NewRouteContext()
	for k, v := range params {
		ctx.URLParams.Add(k, v)
	}

	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
}
