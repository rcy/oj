package u

import (
	"database/sql"
	"html/template"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/gradients"
	"oj/models/users"
	"oj/templatehelpers"

	"github.com/go-chi/chi/v5"
)

var userPageTemplate = template.Must(template.New("layout.html").Funcs(templatehelpers.FuncMap).ParseFiles(layout.File, "handlers/u/user_page.html"))

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

	canEdit := l.User.ID == pageUser.ID

	d := struct {
		Layout  layout.Data
		User    users.User
		CanEdit bool
	}{
		Layout:  l,
		User:    *pageUser,
		CanEdit: canEdit,
	}

	render.Execute(w, userPageTemplate, d)
}
