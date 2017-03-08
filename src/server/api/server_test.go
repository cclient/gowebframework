package api

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"server/common/httputils"
	"testing"
)

func TestMiddlewares(t *testing.T) {
	fmt.Println("test")
	cfg := &Config{}
	srv := &Server{
		cfg: cfg,
	}
	req, _ := http.NewRequest("GET", "/containers/json", nil)
	resp := httptest.NewRecorder()
	ctx := context.Background()
	localHandler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		if httputils.VersionFromContext(ctx) == "" {
			t.Fatalf("Expected version, got empty string")
		}
		return nil
	}
	handlerFunc := srv.handleWithGlobalMiddlewares(localHandler)
	if err := handlerFunc(ctx, resp, req, map[string]string{}); err != nil {
		t.Fatal(err)
	}
}
