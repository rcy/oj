package bots

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/render"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func provideBot(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		query := api.New(db.DB)

		botID, _ := strconv.Atoi(chi.URLParam(r, "botID"))
		bot, err := query.Bot(ctx, int64(botID))
		if errors.Is(err, sql.ErrNoRows) {
			render.NotFound(w)
			return
		}
		if err != nil {
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ctx = context.WithValue(ctx, "bot", bot)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func botFromContext(ctx context.Context) api.Bot {
	return ctx.Value("bot").(api.Bot)
}
