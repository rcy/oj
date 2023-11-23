package layout

import (
	"context"
	"net/http"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
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

func NewContext(ctx context.Context, data Data) context.Context {
	return context.WithValue(ctx, layoutContextKey, data)
}

func Provider(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := auth.FromContext(ctx)
		data, err := FromUser(user)
		if err != nil {
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ctx = NewContext(ctx, data)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
