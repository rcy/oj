package create

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"

	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Get("/", page)
	r.Post("/", post)
}

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent, pageContent)
)

func page(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())

	render.Execute(w, pageTemplate, struct {
		Layout layout.Data
	}{
		Layout: l,
	})
}

func post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queries := api.New(db.DB)

	quiz, err := queries.CreateQuiz(ctx, api.CreateQuizParams{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	})
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/admin/quizzes/%d", quiz.ID), http.StatusSeeOther)
}
