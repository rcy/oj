package become

import (
	"context"
	"net/http"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"oj/models/users"
)

func Provider(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		becomeUser, err := getUser(ctx)
		if err != nil {
			if err == auth.ErrNotAuthorized {
				render.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if becomeUser != nil {
			ctx = auth.NewContext(ctx, *becomeUser)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func getUser(ctx context.Context) (*users.User, error) {
	user := auth.FromContext(ctx)
	if user.BecomeUserID == nil {
		return nil, nil
	}
	if !user.Admin {
		return nil, auth.ErrNotAuthorized
	}
	return users.FindById(*user.BecomeUserID)
}
