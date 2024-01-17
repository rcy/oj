package handlers

import (
	"log"
	"net/http"
	"oj/handlers/admin"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/handlers/welcome"
	"oj/internal/app"
	"oj/internal/middleware/auth"
	"oj/internal/middleware/become"
	"oj/internal/middleware/redirect"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
)

func Router(db *sqlx.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)

	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags), NoColor: true})
	r.Use(middleware.Logger)

	//r.Use(oj.Redirect)

	// authenticated routes
	r.Route("/", func(r chi.Router) {
		r.Use(auth.Provider)
		r.Use(become.Provider)
		r.Use(redirect.Redirect)
		r.Use(layout.Provider)
		r.Handle("/*", app.Handler(db))
	})

	r.Route("/admin", func(r chi.Router) {
		r.Use(auth.Provider)
		r.Use(layout.Provider)
		r.Handle("/*", admin.Handler(db))
	})

	// non authenticated routes
	r.Route("/welcome", welcome.Route)

	// serve static files
	fs := http.FileServer(http.Dir("assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets", fs))

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/favicon.ico")
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.Error(w, "Page not found", 404)
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		render.Error(w, "Method not allowed", 405)
	})

	return r
}
