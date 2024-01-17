package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	router := Router(nil)

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
