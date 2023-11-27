package postoffice

import (
	_ "embed"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/postoffice/compose"
	"oj/handlers/render"

	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/postoffice/inbox", http.StatusSeeOther)
	})

	r.Get("/inbox", page)
	r.Route("/compose", compose.Router)
}

var (
	//go:embed page.gohtml
	pageContent string

	pageTemplate = layout.MustParse(pageContent)
)

func page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queries := api.New(db.DB)
	l := layout.FromContext(r.Context())

	received, err := queries.UserPostcardsReceived(ctx, l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sent, err := queries.UserPostcardsSent(ctx, l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout   layout.Data
		Received []api.UserPostcardsReceivedRow
		Sent     []api.UserPostcardsSentRow
	}{
		Layout:   l,
		Received: received,
		Sent:     sent,
	})
}
