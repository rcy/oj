package compose

import (
	_ "embed"
	"log"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Get("/", page)
	r.Post("/", post)
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

	connections, err := queries.GetConnections(ctx, l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	render.Execute(w, pageTemplate, struct {
		Layout      layout.Data
		Connections []api.User
	}{
		Layout:      l,
		Connections: connections,
	})
}

func post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queries := api.New(db.DB)
	sender := auth.FromContext(ctx)

	recipient, _ := strconv.Atoi(r.FormValue("recipient"))
	params := api.CreatePostcardParams{
		Sender:    sender.ID,
		Recipient: int64(recipient),
		Subject:   r.FormValue("subject"),
		Body:      r.FormValue("body"),
		State:     "queued",
	}

	log.Print("postcard", params)

	_, err := queries.CreatePostcard(ctx, params)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/postoffice", http.StatusSeeOther)
}
