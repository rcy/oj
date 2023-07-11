package connect

import (
	"html/template"
	"net/http"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/users"

	"github.com/go-chi/chi/v5"
)

var t = template.Must(template.ParseFiles(layout.File, "handlers/connect/connect.html"))

type Connection struct {
	users.User
	RoleIn  *string `db:"role_in"`
	RoleOut *string `db:"role_out"`
}

func GetConnection(user1ID int64, user2ID any) (*Connection, error) {
	var connection Connection
	err := db.DB.Get(&connection, `
select u.*,
       case
           when f1.a_id = $1 then f1.b_role
           else null
       end as role_out,
       case
           when f2.b_id = $1 then f2.b_role
           else null
       end as role_in
from users u
left join friends f1 on f1.b_id = u.id and f1.a_id = $1
left join friends f2 on f2.a_id = u.id and f2.b_id = $1
where
  u.id = $2
`, user1ID, user2ID)
	if err != nil {
		return nil, err
	}
	return &connection, err
}

func (f Connection) Status() string {
	if f.RoleOut == nil {
		if f.RoleIn == nil {
			return "none"
		} else {
			return "request received"
		}
	} else {
		if f.RoleIn == nil {
			return "request sent"
		} else {
			return "connected"
		}
	}
}

func Connect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	lay, _ := layout.FromContext(ctx)

	var connections []Connection
	err := db.DB.Select(&connections, `
select u.*,
       case
           when f1.a_id = $1 then f1.b_role
           else null
       end as role_out,
       case
           when f2.b_id = $1 then f2.b_role
           else null
       end as role_in
from users u
left join friends f1 on f1.b_id = u.id and f1.a_id = $1
left join friends f2 on f2.a_id = u.id and f2.b_id = $1
where
  u.id != $1
and
  is_parent = 1
order by role_in desc
limit 128;
`, lay.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, t, "layout.html", struct {
		Layout      layout.Data
		Connections []Connection
	}{
		Layout:      lay,
		Connections: connections,
	})
}

func PutParentFriend(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := users.FromContext(ctx)
	userID := chi.URLParam(r, "userID")

	var user users.User
	err := db.DB.Get(&user, `select * from users where id = $1`, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if !user.IsParent {
		http.Error(w, "not a parent", http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec(`insert into friends(a_id, b_id, b_role) values(?,?,'friend')`, currentUser.ID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	connection, err := GetConnection(currentUser.ID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, t, "connection", connection)
}

func DeleteParentFriend(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := users.FromContext(ctx)
	userID := chi.URLParam(r, "userID")

	var user users.User
	err := db.DB.Get(&user, `select * from users where id = $1`, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if !user.IsParent {
		http.Error(w, "not a parent", http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec(`delete from friends where a_id = $1 and b_id = $2`, currentUser.ID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	connection, err := GetConnection(currentUser.ID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, t, "connection", connection)
}
