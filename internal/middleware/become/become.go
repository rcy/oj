package become

import (
	"net/http"
	"oj/handlers/render"
	"oj/models/users"
)

func Provider(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		becomeUser, err := users.BecomeFromContext(ctx)
		if err != nil {
			if err == users.ErrNotAuthorized {
				render.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if becomeUser != nil {
			ctx = users.NewContext(ctx, *becomeUser)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
