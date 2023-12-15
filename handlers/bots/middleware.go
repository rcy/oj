package bots

import (
	"context"
	"net/http"
	"oj/handlers/bots/ai"
	"oj/handlers/render"

	"github.com/go-chi/chi/v5"
	"github.com/sashabaranov/go-openai"
)

func provideAssistant(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		assistant, err := ai.New().Client.RetrieveAssistant(ctx, chi.URLParam(r, "assistantID"))
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
