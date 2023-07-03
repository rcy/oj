package chat

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"oj/db"
	"oj/handlers/eventsource"
	"oj/handlers/layout"
	"oj/handlers/render"
	"sync"

	"oj/models/gradients"
	"oj/models/messages"
	"oj/models/rooms"
	"oj/models/users"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

var chatTemplate = template.Must(
	template.ParseFiles(
		layout.File,
		"handlers/chat/chat_index_ws.html",
	))

func UserChatPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := users.FromContext(ctx)

	pageUserID := chi.URLParam(r, "userID")
	pageUser, err := users.FindByStringId(pageUserID)
	if err != nil {
		if err == sql.ErrNoRows {
			render.Error(w, "User not found", 404)
			return
		}
	}
	ug, err := gradients.UserBackground(pageUser.ID)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	room, err := rooms.FindOrCreateByUserIDs(user.ID, pageUser.ID)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	var records []messages.Message
	err = db.DB.Select(&records, `
select * from
(
 select m.*, sender.avatar_url as sender_avatar_url from messages m
   join users sender on m.sender_id = sender.id
   where m.room_id = ?
   order by created_at desc
   limit 128
)
order by created_at asc
`, room.ID, user.ID)
	if err != nil {
		render.Error(w, "selecting messages: "+err.Error(), 500)
		return
	}

	err = updateDeliveries(db.DB, room.ID, user.ID)
	if err != nil {
		render.Error(w, "marking deliveries sent: "+err.Error(), 500)
		return
	}

	// get the layout after the deliveries have been updated
	l, err := layout.GetData(r)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}
	// override layout gradient to show the page user's not the request user's
	l.BackgroundGradient = *ug

	pd := struct {
		Layout   layout.Data
		User     users.User
		RoomID   int64
		Messages []messages.Message
	}{
		Layout:   l,
		User:     *pageUser,
		RoomID:   room.ID,
		Messages: records,
	}

	render.Execute(w, chatTemplate, pd)
}

var udMut sync.Mutex

func updateDeliveries(db *sqlx.DB, roomID, userID int64) error {
	udMut.Lock()
	defer udMut.Unlock()

	log.Printf("UPDATE DELIVERIES %d", userID)
	_, err := db.DB.Exec(`update deliveries set sent_at = current_timestamp where sent_at is null and room_id = ? and recipient_id = ?`, roomID, userID)
	log.Printf("UPDATE DELIVERIES %d...done", userID)

	eventsource.SSE.SendMessage(
		fmt.Sprintf("/es/user-%d", userID),
		sse.NewMessage("", "simple", "USER_UPDATE"))

	return err
}
