package u

import (
	"html/template"
	"net/http"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/users"
)

var uit = template.Must(template.ParseFiles(layout.File, "handlers/u/people.html"))

type UserWithCount struct {
	users.User
	Count int
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)

	l, err := layout.GetData(r)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	var allUsers []UserWithCount

	err = db.DB.Select(&allUsers, "select * from users where id <> ?", user.ID)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	for i, recipient := range allUsers {
		err := db.DB.Get(&allUsers[i].Count, `select count() from deliveries where recipient_id = ? and sender_id = ? and sent_at is null`, user.ID, recipient.ID)
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
		AllUsers: allUsers,
	}

	render.Execute(w, uit, d)
}
