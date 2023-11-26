package friends

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

	MyPageTemplate = layout.MustParse(pageContent, me.CardContent)
)

func Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := layout.FromContext(r.Context())
	queries := api.New(db.DB)

	friends, err := queries.GetFriendsWithGradient(ctx, l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d := struct {
		Layout  layout.Data
		User    api.User
		Friends []api.GetFriendsWithGradientRow
	}{
		Layout:  l,
		User:    l.User,
		Friends: friends,
	}

	render.Execute(w, MyPageTemplate, d)
}
