package connectkids

import (
	_ "embed"
	"log"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
	"oj/models/users"
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
	lay := layout.FromContext(ctx)
	queries := api.New(db.DB)

	// get possible friend connections
	// find all the kids of all the friends of this kid's parenst

	reachableKids := map[int64]any{}

	parents, err := queries.GetParents(ctx, lay.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, parent := range parents {
		kids, err := queries.GetKids(ctx, parent.ID)
		if err != nil {
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, kids := range kids {
			reachableKids[kids.ID] = true
		}

		friends, err := queries.GetFriends(ctx, parent.ID)
		if err != nil {
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, friend := range friends {
			kids, err := queries.GetKids(ctx, friend.ID)
			if err != nil {
				render.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			for _, kids := range kids {
				reachableKids[kids.ID] = true
			}
		}
	}

	delete(reachableKids, lay.User.ID)

	var connections []api.GetConnectionRow
	for kidID := range reachableKids {
		connection, err := queries.GetConnection(ctx, api.GetConnectionParams{AID: lay.User.ID, ID: kidID})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("%v", connection)
		connections = append(connections, connection)
	}

	render.Execute(w, t, struct {
		Layout      layout.Data
		Connections []api.GetConnectionRow
	}{
		Layout:      lay,
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
