package connectkids

import (
	_ "embed"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"oj/models/users"
	"oj/services/reachable"
	"oj/worker"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var (
	//go:embed connectkids.gohtml
	pageContent string
	t           = layout.MustParse(pageContent)
)

func KidConnect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(ctx)

	connections, err := reachable.ReachableKids(ctx, l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, t, struct {
		Layout      layout.Data
		Connections []api.GetConnectionRow
	}{
		Layout:      l,
		Connections: connections,
	})
}

func PutKidFriend(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := auth.FromContext(ctx)
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))
	queries := api.New(db.DB)

	var user users.User
	err := db.DB.Get(&user, `select * from users where id = $1`, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	result, err := db.DB.Exec(`insert into friends(a_id, b_id, b_role) values(?,?,'friend')`, currentUser.ID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	friendID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	go worker.NotifyKidFriend(friendID)

	connection, err := queries.GetConnection(ctx, api.GetConnectionParams{AID: currentUser.ID, ID: int64(userID)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, t, "connection", connection)
}

func DeleteKidFriend(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := auth.FromContext(ctx)
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))
	queries := api.New(db.DB)

	var user users.User
	err := db.DB.Get(&user, `select * from users where id = $1`, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, err = db.DB.Exec(`delete from friends where a_id = $1 and b_id = $2`, currentUser.ID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	connection, err := queries.GetConnection(ctx, api.GetConnectionParams{AID: currentUser.ID, ID: int64(userID)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, t, "connection", connection)
}
