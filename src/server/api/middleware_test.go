package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	//"errcode"
	//"errors"
	"golang.org/x/net/context"
	"server/common/httputils"
)

//
func TestVersionMiddleware(t *testing.T) {
	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		if httputils.VersionFromContext(ctx) == "" {
			t.Fatalf("Expected version, got empty string")
		}
		return nil
	}
	h := ParseFormMiddleware(handler)
	req, _ := http.NewRequest("GET", "/containers/json", nil)
	resp := httptest.NewRecorder()
	ctx := context.Background()
	if err := h(ctx, resp, req, map[string]string{}); err != nil {
		t.Fatal(err)
	}
}
