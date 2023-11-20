package quizzes

import (
	"database/sql"
	_ "embed"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/admin/quizzes/create"
	"oj/handlers/admin/quizzes/show"
	"oj/handlers/layout"
	"oj/handlers/render"

	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Get("/", page)
	r.Route("/create", create.Router)
	r.Route("/{quizID}", show.Router)
}

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent, pageContent)
)

func page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)

	queries := api.New(db.DB)

	allQuizzes, err := queries.AllQuizzes(ctx)
	if err != nil && err != sql.ErrNoRows {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout  layout.Data
		Quizzes []api.Quiz
	}{
		Layout:  l,
		Quizzes: allQuizzes,
	})
}
