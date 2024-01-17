package app

import (
	"net/http"
	"net/http/httptest"
	"oj/db"
	"testing"
)

func TestApplicationRouter(t *testing.T) {
	handler := Handler(db.DB)

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusFound {
		t.Errorf("unexpected status code, got %d", resp.StatusCode)
	}
}
