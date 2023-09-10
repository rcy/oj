package me

import (
	"html/template"
	"net/http"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/users"
	"oj/templatehelpers"
)

var MyPageTemplate = template.Must(template.New("layout.gohtml").Funcs(templatehelpers.FuncMap).ParseFiles(layout.File, "handlers/me/my_page.gohtml", "handlers/me/card.gohtml"))

type UnreadUser struct {
	users.User
	UnreadCount int `db:"unread_count"`
}

func (uu UnreadUser) GradientCSS() template.CSS {
	return template.CSS("red")
}

func (uu UnreadUser) Role() string {
	return "foo"
}

func MyPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l, err := layout.FromContext(ctx)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	query := `
select users.*, count(*) unread_count
from deliveries
join users on sender_id = users.id
where recipient_id = ? and sent_at is null
group by users.username;
`
	var unreadUsers []UnreadUser
	err = db.DB.Select(&unreadUsers, query, l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d := struct {
		Layout      layout.Data
		User        users.User
		UnreadUsers []UnreadUser
	}{
		Layout:      l,
		User:        l.User,
		UnreadUsers: unreadUsers,
	}

	render.Execute(w, MyPageTemplate, d)
}
