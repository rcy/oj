package me

import (
	_ "embed"
	"html/template"
	"net/http"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/users"
)

var (
	//go:embed card.gohtml
	CardContent string

	//go:embed page.gohtml
	pageContent string

	pageTemplate = layout.MustParse(pageContent, CardContent)
)

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

func Page(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())

	query := `
select users.*, count(*) unread_count
from deliveries
join users on sender_id = users.id
where recipient_id = ? and sent_at is null
group by users.username;
`
	var unreadUsers []UnreadUser
	err := db.DB.Select(&unreadUsers, query, l.User.ID)
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

	render.Execute(w, pageTemplate, d)
}
