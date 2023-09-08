package middleware

import (
	"log"
	"net/http"
	"oj/models/users"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("kh_session")
		if err != nil {
			http.Redirect(w, r, "/welcome?error1", http.StatusSeeOther)
			return
		}

		user, err := users.FromSessionKey(cookie.Value)
		if err != nil {
			log.Printf("error: FromSessionKey %s", err)
			http.Redirect(w, r, "/welcome?error2", http.StatusSeeOther)
			return
		}

		ctx := users.NewContext(r.Context(), user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
