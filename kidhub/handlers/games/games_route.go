package games

import (
	"html/template"
	"net/http"
	"oj/handlers"
	"oj/models/users"

	"github.com/go-chi/chi/v5"
)

func Route(r chi.Router) {
	r.Get("/", index)
}

var t = template.Must(template.ParseFiles("handlers/layout.html", "handlers/games/games_index.html"))

func index(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)
	err := t.Execute(w, struct{ User users.User }{User: user})
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}
}
