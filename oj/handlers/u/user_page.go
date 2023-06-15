package u

import (
	"database/sql"
	"html/template"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/gradients"
	"oj/models/users"

	"github.com/go-chi/chi/v5"
)

var t = template.Must(template.ParseFiles(layout.File, "handlers/u/user_page.html"))

func UserPage(w http.ResponseWriter, r *http.Request) {
	l, err := layout.GetData(r)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	username := chi.URLParam(r, "username")
	user, err := users.FindByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			render.Error(w, "User not found", 404)
			return
		}
	}
	ug, err := gradients.UserBackground(user.ID)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}
	// override layout gradient to show the page user's not the request user's
	l.BackgroundGradient = *ug

	d := struct {
		Layout layout.Data
		User   users.User
	}{
		Layout: l,
		User:   *user,
	}

	render.Execute(w, t, d)
}
