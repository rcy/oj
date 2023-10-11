package quizzes

import (
	"database/sql"
	_ "embed"
	"net/http"
	"oj/handlers/admin/quizzes/show"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/quizzes"

	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Get("/", page)
	r.Route("/{quizID}", show.Router)
}

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent, pageContent)
)

func page(w http.ResponseWriter, r *http.Request) {
	l, err := layout.FromRequest(r)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	allQuizzes, err := quizzes.FindAll()
	if err != nil && err != sql.ErrNoRows {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout  layout.Data
		Quizzes []quizzes.Quiz
	}{
		Layout:  l,
		Quizzes: allQuizzes,
	})
}
