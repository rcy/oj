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

type Unread struct {
	SenderID int64 `db:"sender_id"`
	Count    int
}

type Friend struct {
	users.User
	Role        string
	UnreadCount int
	GradientCSS template.CSS
}

func MyPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l, err := layout.FromContext(ctx)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d := struct {
		Layout layout.Data
		User   users.User
	}{
		Layout: l,
		User:   l.User,
	}

	render.Execute(w, MyPageTemplate, d)
}

func getUnreads(userID int64) ([]Unread, error) {
	var unreads []Unread

	err := db.DB.Select(&unreads, `
	  select sender_id, count(*) count
          from deliveries
          where recipient_id = ? and sent_at is null
          group by sender_id`, userID)
	if err != nil {
		return nil, err
	}
	return unreads, nil
}
