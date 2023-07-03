package friends

import (
	"html/template"
	"net/http"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/users"
)

var uit = template.Must(template.ParseFiles(layout.File, "handlers/friends/friends.html"))

type UserWithCount struct {
	users.User
	Count int
	Role  string
}

func Friends(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := users.FromContext(ctx)

	l, err := layout.GetData(r)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	var friends []UserWithCount

	err = db.DB.Select(&friends, `select users.*, b_role role from friends join users on b_id = users.id where a_id = ?`, user.ID)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	for i, recipient := range friends {
		err := db.DB.Get(&friends[i].Count, `select count() from deliveries where recipient_id = ? and sender_id = ? and sent_at is null`, user.ID, recipient.ID)
		if err != nil {
			render.Error(w, err.Error(), 500)
			return
		}
	}

	d := struct {
		Layout   layout.Data
		User     users.User
		AllUsers []UserWithCount
	}{
		Layout:   l,
		User:     l.User,
		AllUsers: friends,
	}

	render.Execute(w, uit, d)
}
