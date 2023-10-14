package quizzes

import (
	"database/sql"
	_ "embed"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/quizzes"
)

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent)
)

func Page(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())

	allQuizzes, err := quizzes.FindAllPublished()
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
