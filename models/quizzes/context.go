package quizzes

import (
	"context"
	"database/sql"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/render"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type contextKey int

const quizContextKey contextKey = iota

func FromContext(ctx context.Context) api.Quiz {
	return ctx.Value(quizContextKey).(api.Quiz)
}

func Provider(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		queries := api.New(db.DB)
		quizID, _ := strconv.Atoi(chi.URLParam(r, "quizID"))
		quiz, err := queries.Quiz(ctx, int64(quizID))
		if err != nil {
			if err == sql.ErrNoRows {
				render.NotFound(w)
				return
			}
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ctx = context.WithValue(ctx, quizContextKey, quiz)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
