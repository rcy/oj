package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestRouter(t *testing.T) {
	router := Router()

	req := httptest.NewRequest("GET", "/welcome", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
	//body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
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
