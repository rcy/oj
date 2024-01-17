package app

import (
	"net/http"
	"net/http/httptest"
	"oj/db"
	"testing"
)

func TestApplicationRouter(t *testing.T) {
	routes := Resource{DB: db.DB}.Routes()

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	routes.ServeHTTP(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusFound {
		t.Errorf("unexpected status code, got %d", resp.StatusCode)
	}
}
