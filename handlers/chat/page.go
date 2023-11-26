package chat

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/eventsource"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"oj/services/background"
	"oj/services/room"
	"strconv"
	"sync"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent)
)

func Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queries := api.New(db.DB)
	user := auth.FromContext(ctx)

	pageUserID, _ := strconv.Atoi(chi.URLParam(r, "userID"))
	pageUser, err := queries.UserByID(ctx, int64(pageUserID))
	if err != nil {
		if err == sql.ErrNoRows {
			render.Error(w, "User not found", 404)
			return
		}
	}
	ug, err := background.ForUser(ctx, pageUser.ID)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	room, err := room.FindOrCreateByUserIDs(ctx, user.ID, pageUser.ID)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	records, err := queries.RecentMessages(ctx, fmt.Sprint(room.ID))
	if err != nil {
		render.Error(w, "api selecting messages: "+err.Error(), 500)
		return
	}

	err = updateDeliveries(db.DB, room.ID, user.ID)
	if err != nil {
		render.Error(w, "marking deliveries sent: "+err.Error(), 500)
		return
	}

	// get the layout after the deliveries have been updated to ensure unread count is correct
	l, err := layout.FromUser(ctx, user)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}
	// override layout gradient to show the page user's not the request user's
	l.BackgroundGradient = *ug

	pd := struct {
		Layout   layout.Data
		User     api.User
		RoomID   int64
		Messages []api.RecentMessagesRow
	}{
		Layout:   l,
		User:     pageUser,
		RoomID:   room.ID,
		Messages: records,
	}

	render.Execute(w, pageTemplate, pd)
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
