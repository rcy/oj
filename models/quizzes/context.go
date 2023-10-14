package quizzes

import (
	"context"
	"database/sql"
	"net/http"
	"oj/handlers/render"

	"github.com/go-chi/chi/v5"
)

type contextKey int

const quizContextKey contextKey = iota

func FromContext(ctx context.Context) Quiz {
	return ctx.Value(quizContextKey).(Quiz)
}

func Provider(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		quiz, err := FindByStringID(chi.URLParam(r, "quizID"))
		if err != nil {
			if err == sql.ErrNoRows {
				render.NotFound(w)
				return
			}
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ctx = context.WithValue(ctx, quizContextKey, *quiz)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
