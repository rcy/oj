package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"oj/db"
	"oj/handlers/eventsource"
	"oj/handlers/render"
	"oj/models/users"
	"oj/worker"
	"strconv"
	"strings"
	"time"

	"github.com/alexandrevicenzi/go-sse"
)

func PostChatMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := users.FromContext(ctx)
	roomID, err := strconv.Atoi(r.FormValue("roomID"))
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}
	body := r.FormValue("body")

	if strings.TrimSpace(body) != "" {
		err = postMessage(r.Context(), int64(roomID), user.ID, body)
		if err != nil {
			render.Error(w, err.Error(), 500)
			return
		}
	}

	err = chatTemplate.ExecuteTemplate(w, "chatInput", struct{ RoomID int }{RoomID: roomID})
	if err != nil {
		render.Error(w, err.Error(), 500)
	}
}

type RoomUser struct {
	ID        int64
	CreatedAt time.Time `db:"created_at"`
	RoomID    int64     `db:"room_id"`
	UserID    int64     `db:"user_id"`
	Email     *string   `db:"email"`
}

func postMessage(ctx context.Context, roomID, senderID int64, body string) error {
	var roomUsers []RoomUser

	tx, err := db.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var sender users.User

	err = tx.Get(&sender, `select * from users where id = ?`, senderID)
	if err != nil {
		return err
	}

	// get the users of the room
	err = tx.Select(&roomUsers, `select room_users.*, users.email from room_users join users on room_users.user_id = users.id where room_id = ?`, roomID)
	if err != nil {
		return err
	}

	// create the message
	result, err := tx.Exec(`insert into messages(room_id, sender_id, body) values(?,?,?)`, roomID, senderID, body)
	if err != nil {
		return err
	}
	messageID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// create deliveries for each user in the room
	var deliveryIDs []int64
	for _, roomUser := range roomUsers {
		result, err = tx.Exec(`insert into deliveries(message_id, room_id, sender_id, recipient_id) values(?,?,?,?)`, messageID, roomID, senderID, roomUser.UserID)
		if err != nil {
			return err
		}

		deliveryID, err := result.LastInsertId()
		if err != nil {
			return err
		}

		deliveryIDs = append(deliveryIDs, deliveryID)
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	// send notifications after the transaction has been committed

	for _, deliveryID := range deliveryIDs {
		worker.NotifyDelivery(deliveryID)
	}

	data, err := json.Marshal(map[string]interface{}{
		"senderID": fmt.Sprint(senderID),
	})
	if err != nil {
		return err
	}

	eventsource.SSE.SendMessage(
		fmt.Sprintf("/es/room-%d", roomID),
		sse.NewMessage("", string(data), "NEW_MESSAGE"))

	for _, roomUser := range roomUsers {
		if roomUser.UserID == senderID {
			continue
		}

		eventsource.SSE.SendMessage(
			fmt.Sprintf("/es/user-%d", roomUser.UserID),
			sse.NewMessage("", "simple", "USER_UPDATE"))
	}

	return nil
}
