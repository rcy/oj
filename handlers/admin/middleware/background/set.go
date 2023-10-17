package background

import (
	"net/http"
	"oj/element/gradient"
	"oj/handlers/layout"
)

// Returns middleware to override background gradient established by layout middleware
func Set(grad gradient.Gradient) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			l := layout.FromContext(ctx)
			l.BackgroundGradient = grad
			ctx = layout.NewContext(ctx, l)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
