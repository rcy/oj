package me

import (
	_ "embed"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
)

var (
	//go:embed card.gohtml
	CardContent string

	//go:embed page.gohtml
	pageContent string

	pageTemplate = layout.MustParse(pageContent, CardContent)
)

func Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queries := api.New(db.DB)

	l := layout.FromContext(r.Context())

	unreadUsers, err := queries.UsersWithUnreadCounts(ctx, l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d := struct {
		Layout      layout.Data
		User        api.User
		UnreadUsers []api.UsersWithUnreadCountsRow
	}{
		Layout:      l,
		User:        l.User,
		UnreadUsers: unreadUsers,
	}

	render.Execute(w, pageTemplate, d)
}
