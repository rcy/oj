package chat

import (
	"html/template"
	"net/http"
	"oj/handlers"

	"github.com/go-chi/chi/v5"
)

var chatTemplate = template.Must(template.ParseFiles("handlers/layout.html", "handlers/chat/chat_index.html"))

func Route(r chi.Router) {
	r.Get("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	err := chatTemplate.Execute(w, struct{ Username string }{Username: r.Context().Value("username").(string)})
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}
}
