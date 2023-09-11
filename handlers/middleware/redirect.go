package middleware

import (
	"net/http"
	"time"
)

// Redirect to url stored in redirect cookie
func Redirect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("redirect")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		cookie.Expires = time.Unix(0, 0)
		http.SetCookie(w, cookie)
		http.Redirect(w, r, cookie.Value, http.StatusFound)
	})
}
