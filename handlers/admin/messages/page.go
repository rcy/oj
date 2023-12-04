package messages

import (
	_ "embed"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Get("/", page)
	r.Delete("/{messageID}", deleteMessage)
}

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent, pageContent)
)

func page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queries := api.New(db.DB)

	l := layout.FromContext(r.Context())

	messages, err := queries.AdminRecentMessages(ctx)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout   layout.Data
		Messages []api.AdminRecentMessagesRow
	}{
		Layout:   l,
		Messages: messages,
	})
}

func deleteMessage(w http.ResponseWriter, r *http.Request) {
	messageID, _ := strconv.Atoi(chi.URLParam(r, "messageID"))
	ctx := r.Context()
	queries := api.New(db.DB)

	message, err := queries.AdminDeleteMessage(ctx, int64(messageID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, pageTemplate, "message-row", message)
}
