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

type handler struct {
	router   *chi.Mux
	database *sqlx.DB
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func Handler(database *sqlx.DB) *handler {
	h := &handler{router: chi.NewRouter(), database: database}

	h.router.Use(auth.EnsureAdmin)
	h.router.Use(background.Set(gradient.Admin))
	h.router.Get("/", h.page)
	h.router.Route("/quizzes", quizzes.Router)
	h.router.Route("/messages", messages.Router)
	return h
}

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent, pageContent)
)

func (h *handler) page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queries := api.New(h.database)

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
