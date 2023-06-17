package chat

import (
	"database/sql"
	"html/template"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"

	"oj/models/gradients"
	"oj/models/messages"
	"oj/models/users"

	"github.com/go-chi/chi/v5"
)

var chatTemplate = template.Must(
	template.ParseFiles(
		layout.File,
		"handlers/chat/chat_index_ws.html",
		"handlers/chat/chat_partials.html",
	))

var partials = template.Must(
	template.ParseFiles("handlers/chat/chat_partials.html"),
)

func DM(w http.ResponseWriter, r *http.Request) {
	l, err := layout.GetData(r)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	username := chi.URLParam(r, "username")
	user, err := users.FindByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			render.Error(w, "User not found", 404)
			return
		}
	}
	ug, err := gradients.UserBackground(user.ID)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}
	// override layout gradient to show the page user's not the request user's
	l.BackgroundGradient = *ug

	roomID := users.MakeRoomId(l.User.ID, user.ID)

	records, err := messages.Fetch(roomID)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	pd := struct {
		Layout   layout.Data
		User     users.User
		RoomID   string
		Messages []messages.Message
	}{
		Layout:   l,
		User:     *user,
		RoomID:   roomID,
		Messages: records,
	}

	render.Execute(w, chatTemplate, pd)
}

func PostMessage(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)
	body := r.FormValue("body")

	message, err := messages.Create("dummy-room-id", body, user.Username)
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
