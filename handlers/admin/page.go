package admin

import (
	_ "embed"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/users"
)

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent, pageContent)
)

func Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l, err := layout.FromContext(ctx)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	allUsers, err := users.FindAll()
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Execute(w, pageTemplate, struct {
		Layout layout.Data
		Users  []users.User
	}{
		Layout: l,
		Users:  allUsers,
	})
}
