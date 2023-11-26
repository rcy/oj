package connect

import (
	_ "embed"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"oj/worker"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var (
	//go:embed connect.gohtml
	pageContent string

	//go:embed connection.gohtml
	ConnectionContent string

	t = layout.MustParse(pageContent, ConnectionContent)
)

func Connect(w http.ResponseWriter, r *http.Request) {
	lay := layout.FromContext(r.Context())
	queries := api.New(db.DB)

	connections, err := queries.GetCurrentAndPotentialParentConnections(r.Context(), lay.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, t, struct {
		Layout      layout.Data
		Connections []api.GetCurrentAndPotentialParentConnectionsRow
	}{
		Layout:      lay,
		Connections: connections,
	})
}

func PutParentFriend(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := auth.FromContext(ctx)
	queries := api.New(db.DB)
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))

	user, err := queries.ParentByID(ctx, int64(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	result, err := db.DB.Exec(`insert into friends(a_id, b_id, b_role) values(?,?,'friend')`, currentUser.ID, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	friendID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	go worker.NotifyFriend(friendID)

	connection, err := queries.GetConnection(ctx, api.GetConnectionParams{AID: currentUser.ID, ID: user.ID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("HX-Trigger", "connectionChange")
	render.ExecuteNamed(w, t, "connection", connection)
}

func DeleteParentFriend(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := auth.FromContext(ctx)
	queries := api.New(db.DB)

	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))

	user, err := queries.ParentByID(ctx, int64(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, err = db.DB.Exec(`delete from friends where a_id = $1 and b_id = $2`, currentUser.ID, user.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	connection, err := queries.GetConnection(ctx, api.GetConnectionParams{AID: currentUser.ID, ID: int64(user.ID)})
	if err != nil {
		render.Error(w, "xxx"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("HX-Trigger", "connectionChange")
	render.ExecuteNamed(w, t, "connection", connection)
}
