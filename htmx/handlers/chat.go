package handlers

import (
	"html/template"
	"net/http"
)

var chatTemplate = template.Must(template.ParseFiles("templates/layout.html", "templates/chat.html"))

func Chat(w http.ResponseWriter, r *http.Request) {
	err := chatTemplate.Execute(w, struct{ Username string }{Username: r.Context().Value("username").(string)})
	if err != nil {
		Error(w, err.Error(), 500)
	}
}
