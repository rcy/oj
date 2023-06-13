package chat

import (
	"html/template"
	"net/http"
	"oj/element/gradient"
	"oj/handlers"

	"oj/models/gradients"
	"oj/models/messages"
	"oj/models/users"

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
	user := users.Current(r)

	backgroundGradient, err := gradients.UserBackground(user.ID)
	if err != nil {
		handlers.Error(w, err.Error(), 500)
		return
	}

	records, err := messages.Fetch()
	if err != nil {
		handlers.Error(w, err.Error(), 500)
		return
	}

	pd := struct {
		User               users.User
		BackgroundGradient gradient.Gradient
		Messages           []messages.Message
	}{
		User:               user,
		BackgroundGradient: backgroundGradient,
		Messages:           records,
	}

	err = chatTemplate.Execute(w, pd)
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)
	body := r.FormValue("body")

	message, err := messages.Create(body, user.Username)
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
	user := users.Current(r)
	records, err := messages.Fetch()
	if err != nil {
		handlers.Error(w, err.Error(), 500)
		return
	}

	pd := struct {
		User     users.User
		Messages []messages.Message
	}{
		User:     user,
		Messages: records,
	}

	err = chatTemplate.ExecuteTemplate(w, "chat_messages", pd)
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}
}
