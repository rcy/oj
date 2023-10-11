package admin

import (
	"net/http/httptest"
	"oj/models/users"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestRouter(t *testing.T) {
	router := chi.NewRouter().Route("/", Router)

	for _, tc := range []struct {
		name           string
		path           string
		user           users.User
		wantStatusCode int
	}{
		{
			name:           "non admin user",
			path:           "/",
			user:           users.User{Admin: false},
			wantStatusCode: 401,
		},
		{
			name:           "admin user",
			path:           "/",
			user:           users.User{Admin: true},
			wantStatusCode: 200,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", tc.path, nil)
			ctx := users.NewContext(req.Context(), tc.user)
			router.ServeHTTP(w, req.WithContext(ctx))
			resp := w.Result()

			if resp.StatusCode != tc.wantStatusCode {
				t.Errorf("bad status code, expected %d, got %d", tc.wantStatusCode, resp.StatusCode)
			}
		})
	}
}
