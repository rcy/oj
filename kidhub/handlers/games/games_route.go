package games

import (
	"html/template"
	"net/http"
	"oj/handlers"

	"github.com/go-chi/chi/v5"
)

func Route(r chi.Router) {
	r.Get("/", index)
}

var t = template.Must(template.ParseFiles("handlers/layout.html", "handlers/games/games_index.html"))

func index(w http.ResponseWriter, r *http.Request) {
	err := t.Execute(w, struct{ Username string }{Username: r.Context().Value("username").(string)})
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}
}
