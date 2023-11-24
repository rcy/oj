package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestRouter(t *testing.T) {
	router := Router()

	for _, tc := range []struct {
		name           string
		path           string
		wantStatusCode int
	}{
		{
			name:           "unauthenticated welcome",
			path:           "/welcome",
			wantStatusCode: 200,
		},
		{
			name:           "unauthenticated admin",
			path:           "/admin",
			wantStatusCode: 303,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", tc.path, nil)
			router.ServeHTTP(w, req)
			resp := w.Result()

			if resp.StatusCode != tc.wantStatusCode {
				t.Errorf("bad status code, expected %d, got %d", tc.wantStatusCode, resp.StatusCode)
			}
		})
	}
}

func TestApplicationRouter(t *testing.T) {
	router := chi.NewRouter().Route("/", applicationRouter)

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusFound {
		t.Errorf("unexpected status code, got %d", resp.StatusCode)
	}
}
