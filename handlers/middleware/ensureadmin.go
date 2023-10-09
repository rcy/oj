package middleware

import (
	"net/http"
	"oj/handlers/render"
	"oj/models/users"
)

func EnsureAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := users.FromContext(r.Context())
		if !user.Admin {
			render.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
