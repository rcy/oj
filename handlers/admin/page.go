package admin

import (
	_ "embed"
	"net/http"
	"oj/handlers/admin/quizzes"
	"oj/handlers/layout"
	mw "oj/handlers/middleware"
	"oj/handlers/render"
	"oj/models/users"

	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Use(mw.EnsureAdmin)
	r.Get("/", page)
	r.Route("/quizzes", quizzes.Router)
}

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent, pageContent)
)

func page(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())

	allUsers, err := users.FindAll()
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout layout.Data
		Users  []users.User
	}{
		Layout: l,
		Users:  allUsers,
	})
}
