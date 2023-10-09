package u

import (
	"database/sql"
	_ "embed"
	"net/http"
	"oj/handlers/connect"
	"oj/handlers/layout"
	"oj/handlers/me"
	"oj/handlers/render"
	"oj/models/gradients"
	"oj/models/users"

	"github.com/go-chi/chi/v5"
)

var (
	//go:embed user_page.gohtml
	pageContent string

	userPageTemplate = layout.MustParse(pageContent, me.CardContent, connect.ConnectionContent)
)

func UserPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l, err := layout.FromContext(ctx)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID := chi.URLParam(r, "userID")
	pageUser, err := users.FindByStringId(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			render.Error(w, "User not found", http.StatusNotFound)
			return
		}
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if l.User.ID == pageUser.ID {
		http.Redirect(w, r, "/me", http.StatusFound)
		return
	}

	ug, err := gradients.UserBackground(pageUser.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// override layout gradient to show the page user's not the request user's
	l.BackgroundGradient = *ug

	connection, err := connect.GetConnection(l.User.ID, pageUser.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	canChat := connection.RoleIn != "" && connection.RoleOut != ""

	d := struct {
		Layout     layout.Data
		User       users.User
		Connection *connect.Connection
		CanChat    bool
	}{
		Layout:     l,
		User:       *pageUser,
		Connection: connection,
		CanChat:    canChat,
	}

	render.Execute(w, userPageTemplate, d)
}
