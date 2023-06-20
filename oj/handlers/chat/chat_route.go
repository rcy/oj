package chat

import (
	"database/sql"
	"html/template"
	"net/http"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"

	"oj/models/gradients"
	"oj/models/messages"
	"oj/models/rooms"
	"oj/models/users"

	"github.com/go-chi/chi/v5"
)

var chatTemplate = template.Must(
	template.ParseFiles(
		layout.File,
		"handlers/chat/chat_index_ws.html",
	))

func UserChatPage(w http.ResponseWriter, r *http.Request) {
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

	room, err := rooms.FindOrCreateByUserIDs(l.User.ID, user.ID)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	var records []messages.Message
	err = db.DB.Select(&records, `
select * from
(
 select messages.* from deliveries
   join messages on messages.id = deliveries.message_id
   where messages.room_id = ? and recipient_id = ?
   order by created_at desc
   limit 1000
)
order by created_at asc
`, room.ID, user.ID)
	if err != nil {
		render.Error(w, "selecting deliveries: "+err.Error(), 500)
		return
	}

	// XXX mark stuff as read here

	pd := struct {
		Layout   layout.Data
		User     users.User
		RoomID   int64
		Messages []messages.Message
	}{
		Layout:   l,
		User:     *user,
		RoomID:   room.ID,
		Messages: records,
	}

	render.Execute(w, chatTemplate, pd)
}
