package chat

import (
	"html/template"
	"net/http"
	"oj/handlers"

	"oj/models/messages"

	"github.com/go-chi/chi/v5"
)

var chatTemplate = template.Must(
	template.ParseFiles(
		"handlers/layout.html",
		"handlers/chat/chat_index.html",
		"handlers/chat/chat_partials.html",
	))

var partials = template.Must(
	template.ParseFiles("handlers/chat/chat_partials.html"),
)

func Route(r chi.Router) {
	r.Get("/", index)
	r.Post("/messages", postMessage)
	r.Get("/messages", getMessages)
}

func index(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)

	records, err := messages.Fetch()
	if err != nil {
		handlers.Error(w, err.Error(), 500)
		return
	}

	pd := struct {
		Username string
		Messages []messages.Message
	}{
		Username: username,
		Messages: records,
	}

	err = chatTemplate.Execute(w, pd)
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	body := r.FormValue("body")

	message, err := messages.Create(body, username)
	if err != nil {
		handlers.Error(w, err.Error(), 500)
		return
	}

	w.Header().Add("HX-Trigger", "newMessage")

	err = partials.ExecuteTemplate(w, "chat_input", message)
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	records, err := messages.Fetch()
	if err != nil {
		handlers.Error(w, err.Error(), 500)
		return
	}

	pd := struct {
		Username string
		Messages []messages.Message
	}{
		Username: username,
		Messages: records,
	}

	err = chatTemplate.ExecuteTemplate(w, "chat_messages", pd)
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}
}
