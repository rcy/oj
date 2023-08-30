package friends

import (
	"html/template"
	"net/http"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/gradients"
	"oj/models/users"
)

var uit = template.Must(template.ParseFiles(layout.File, "handlers/friends/friends.gohtml"))

type UserWithCount struct {
	users.User
	Count       int
	Role        string
	GradientCSS template.CSS
}

func Friends(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := users.FromContext(ctx)

	l, err := layout.FromContext(ctx)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	var friends []UserWithCount

	err = db.DB.Select(&friends, `
select users.*, fi.b_role role
from users
join friends fi on fi.b_id = users.id and fi.a_id = $1
join friends fo on fo.a_id = users.id and fo.b_id = $1
`, user.ID)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	for i, recipient := range friends {
		err := db.DB.Get(&friends[i].Count, `
select count()
from deliveries
where
  recipient_id = ?
and
  sender_id = ?
and
  sent_at is null
`, user.ID, recipient.ID)
		if err != nil {
			render.Error(w, err.Error(), 500)
			return
		}

		bg, err := gradients.UserBackground(recipient.ID)
		if err != nil {
			render.Error(w, err.Error(), 500)
			return
		}

		friends[i].GradientCSS = bg.Render()
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
