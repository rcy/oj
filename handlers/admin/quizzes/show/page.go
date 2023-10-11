package show

import (
	"database/sql"
	_ "embed"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/question"
	"oj/models/quizzes"

	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Get("/", page)
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

	q, err := quizzes.FindByStringID(chi.URLParam(r, "quizID"))
	if err != nil {
		if err == sql.ErrNoRows {
			render.NotFound(w)
			return
		}
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	questions, err := q.FindQuestions()
	if err != nil && err != sql.ErrNoRows {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout    layout.Data
		Quiz      *quizzes.Quiz
		Questions []question.Question
	}{
		Layout:    l,
		Quiz:      q,
		Questions: questions,
	})
}
