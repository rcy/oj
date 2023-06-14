package chat

import (
	"html/template"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"

	"oj/models/messages"
	"oj/models/users"
)

var chatTemplate = template.Must(
	template.ParseFiles(
		layout.File,
		"handlers/chat/chat_index.html",
		"handlers/chat/chat_partials.html",
	))

var partials = template.Must(
	template.ParseFiles("handlers/chat/chat_partials.html"),
)

func Index(w http.ResponseWriter, r *http.Request) {
	l, err := layout.GetData(r)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	records, err := messages.Fetch()
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	pd := struct {
		Layout   layout.Data
		User     users.User
		Messages []messages.Message
	}{
		Layout:   l,
		User:     l.User,
		Messages: records,
	}

	render.Execute(w, chatTemplate, pd)
}

func PostMessage(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)
	body := r.FormValue("body")

	message, err := messages.Create(body, user.Username)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	w.Header().Add("HX-Trigger", "newMessage")

	err = partials.ExecuteTemplate(w, "chat_input", message)
	if err != nil {
		render.Error(w, err.Error(), 500)
	}
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)
	records, err := messages.Fetch()
	if err != nil {
		render.Error(w, err.Error(), 500)
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
		render.Error(w, err.Error(), 500)
	}
}
