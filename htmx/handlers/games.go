package handlers

import (
	"html/template"
	"net/http"
)

var gamesTemplate = template.Must(template.ParseFiles("templates/layout.html", "templates/games.html"))

func Games(w http.ResponseWriter, r *http.Request) {
	err := gamesTemplate.Execute(w, struct{ Username string }{Username: r.Context().Value("username").(string)})
	if err != nil {
		Error(w, err.Error(), 500)
	}
}
