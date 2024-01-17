package admin

import (
	_ "embed"
	"net/http"
	"oj/api"
	"oj/element/gradient"
	"oj/handlers/admin/messages"
	"oj/handlers/admin/middleware/auth"
	"oj/handlers/admin/middleware/background"
	"oj/handlers/admin/quizzes"
	"oj/handlers/layout"
	"oj/handlers/render"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type Resource struct {
	DB *sqlx.DB
}

func (rs Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Use(auth.EnsureAdmin)
	r.Use(background.Set(gradient.Admin))
	r.Get("/", rs.page)
	r.Route("/quizzes", quizzes.Router)
	r.Route("/messages", messages.Router)
	return r
}

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent, pageContent)
)

func (rs Resource) page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queries := api.New(rs.DB)

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
