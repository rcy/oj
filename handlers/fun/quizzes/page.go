package quizzes

import (
	"database/sql"
	_ "embed"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
)

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent)
)

func Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)
	queries := api.New(db.DB)

	allQuizzes, err := queries.PublishedQuizzes(ctx)
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
