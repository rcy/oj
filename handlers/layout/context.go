package layout

import (
	"context"
	"net/http"
	"oj/handlers/render"
	"oj/models/users"
)

type contextKey int

const layoutContextKey contextKey = iota

func FromContext(ctx context.Context) Data {
	value := ctx.Value(layoutContextKey)
	if value == nil {
		return Data{}
	}
	return value.(Data)
}

func Provider(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := users.FromContext(ctx)
		data, err := FromUser(user)
		if err != nil {
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ctx = context.WithValue(ctx, layoutContextKey, data)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
