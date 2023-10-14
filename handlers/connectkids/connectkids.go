package connectkids

import (
	"context"
	_ "embed"
	"log"
	"net/http"
	"oj/db"
	"oj/handlers/connect"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/users"
	"oj/worker"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

var (
	//go:embed connectkids.gohtml
	pageContent string
	t           = layout.MustParse(pageContent)
)

func KidConnect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	lay := layout.FromContext(ctx)

	// get possible friend connections
	// find all the kids of all the friends of this kid's parenst

	reachableKids := map[int64]any{}

	parents, err := GetParents(ctx, db.DB, lay.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, parent := range parents {
		kids, err := GetKids(ctx, db.DB, parent.ID)
		if err != nil {
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, kids := range kids {
			reachableKids[kids.ID] = true
		}

		friends, err := GetFriends(ctx, db.DB, parent.ID)
		if err != nil {
			render.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, friend := range friends {
			kids, err := GetKids(ctx, db.DB, friend.ID)
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

	var connections []connect.Connection
	for kidID := range reachableKids {
		connection, err := connect.GetConnection(lay.User.ID, kidID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("%v", connection)
		connections = append(connections, *connection)
	}

	render.Execute(w, t, struct {
		Layout      layout.Data
		Connections []connect.Connection
	}{
		Layout:      lay,
		Connections: connections,
	})
}

func GetParents(ctx context.Context, db *sqlx.DB, userID int64) ([]users.User, error) {
	var parents []users.User
	err := db.Select(&parents, `
		select u.* from users u
		join friends f1 on f1.b_id = u.id and f1.a_id = $1 and f1.b_role = 'parent'
		join friends f2 on f2.a_id = u.id and f2.b_id = $1 and f2.b_role = 'child'
	`, userID)
	if err != nil {
		return nil, err
	}
	return parents, nil
}

func GetKids(ctx context.Context, db *sqlx.DB, userID int64) ([]users.User, error) {
	var kids []users.User
	err := db.Select(&kids, `
		select u.* from users u
		join friends f1 on f1.b_id = u.id and f1.a_id = $1 and f1.b_role = 'child'
		join friends f2 on f2.a_id = u.id and f2.b_id = $1 and f2.b_role = 'parent'
	`, userID)
	if err != nil {
		return nil, err
	}
	return kids, nil
}

func GetFriends(ctx context.Context, db *sqlx.DB, userID int64) ([]users.User, error) {
	var friends []users.User
	err := db.Select(&friends, `
		select u.* from users u
		join friends f1 on f1.b_id = u.id and f1.a_id = $1 and f1.b_role = 'friend'
		join friends f2 on f2.a_id = u.id and f2.b_id = $1 and f2.b_role = 'friend'
	`, userID)
	if err != nil {
		return nil, err
	}
	return friends, nil
}

func PutKidFriend(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := users.FromContext(ctx)
	userID := chi.URLParam(r, "userID")

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

	connection, err := connect.GetConnection(currentUser.ID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, t, "connection", connection)
}

func DeleteKidFriend(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := users.FromContext(ctx)
	userID := chi.URLParam(r, "userID")

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

	connection, err := connect.GetConnection(currentUser.ID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, t, "connection", connection)
}
