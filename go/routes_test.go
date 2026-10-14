package main

import (
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"testing"
)

func get(t *testing.T, path string) *httptest.ResponseRecorder {
	t.Helper()

	recorder := httptest.NewRecorder()
	routes().ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, path, nil))

	if recorder.Code != http.StatusOK {
		t.Fatalf("GET %s: expected status 200, got %d", path, recorder.Code)
	}

	return recorder
}

func TestWelcomePageResponds(t *testing.T) {
	body := get(t, "/").Body.String()

	for _, want := range []string{"Your Go app is up and running on Laravel Cloud.", "Go " + strings.TrimPrefix(runtime.Version(), "go")} {
		if !strings.Contains(body, want) {
			t.Fatalf("expected the welcome page to contain %q", want)
		}
	}
}

func TestAPIResponds(t *testing.T) {
	if body := get(t, "/api").Body.String(); !strings.Contains(body, "Hello from Laravel Cloud") {
		t.Fatalf("unexpected body: %s", body)
	}
}
