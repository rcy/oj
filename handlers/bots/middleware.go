package bots

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/bots/ai"
	"oj/handlers/render"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sashabaranov/go-openai"
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

func provideAssistant(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		bot := botFromContext(ctx)

		assistant, err := ai.New().Client.RetrieveAssistant(ctx, bot.AssistantID)
		if err != nil {
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ctx = context.WithValue(ctx, "assistant", assistant)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func assistantFromContext(ctx context.Context) openai.Assistant {
	return ctx.Value("assistant").(openai.Assistant)
}
