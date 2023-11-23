package become

import (
	"context"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
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

func getUser(ctx context.Context) (*api.User, error) {
	queries := api.New(db.DB)
	user := auth.FromContext(ctx)
	if !user.BecomeUserID.Valid {
		return nil, nil
	}
	if !user.Admin {
		return nil, auth.ErrNotAuthorized
	}
	becomeUser, err := queries.UserByID(ctx, user.BecomeUserID.Int64)
	return &becomeUser, err
}
