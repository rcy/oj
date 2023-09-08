package middleware

import (
	"log"
	"net/http"
	"oj/handlers/render"
	"oj/models/users"
)

func Become(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		becomeUser, err := users.Become(ctx)
		if err != nil {
			log.Printf("error: Become %s", err)
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if becomeUser != nil {
			user := users.FromContext(ctx)
			log.Printf("*** user %d becoming user %d", user.ID, becomeUser.ID)
			ctx = users.NewContext(ctx, *becomeUser)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
