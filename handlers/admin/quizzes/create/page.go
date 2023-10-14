package create

import (
	_ "embed"
	"fmt"
	"net/http"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/quizzes"

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
	quiz := quizzes.Quiz{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	}
	result, err := db.DB.Exec(`insert into quizzes(name,description) values(?,?)`, quiz.Name, quiz.Description)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	quizID, err := result.LastInsertId()
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/admin/quizzes/%d", quizID), http.StatusSeeOther)
}
