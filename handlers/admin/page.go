package admin

import (
	_ "embed"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/element/gradient"
	"oj/handlers/admin/messages"
	"oj/handlers/admin/middleware/auth"
	"oj/handlers/admin/middleware/background"
	"oj/handlers/admin/quizzes"
	"oj/handlers/layout"
	"oj/handlers/render"

	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Use(auth.EnsureAdmin)
	r.Use(background.Set(gradient.Admin))
	r.Get("/", page)
	r.Route("/quizzes", quizzes.Router)
	r.Route("/messages", messages.Router)
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

	allUsers, err := queries.AllUsers(ctx)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout layout.Data
		Users  []api.User
	}{
		Layout: l,
		Users:  allUsers,
	})
}
