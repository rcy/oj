package humans

import (
	_ "embed"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/me"
	"oj/handlers/render"
)

var (
	//go:embed "page.gohtml"
	pageContent string

	pageTemplate = layout.MustParse(pageContent, me.CardContent)
)

func Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(r.Context())
	queries := api.New(db.DB)

	friends, err := queries.GetConnectionsWithGradient(ctx, l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d := struct {
		Layout  layout.Data
		User    api.User
		Friends []api.GetConnectionsWithGradientRow
	}{
		Layout:  l,
		User:    l.User,
		Friends: friends,
	}

	render.Execute(w, pageTemplate, d)
}
