package editme

import (
	"html/template"
	"net/http"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/users"
	"oj/templatehelpers"
)

var myPageEditTemplate = template.Must(template.New("layout.gohtml").Funcs(templatehelpers.FuncMap).ParseFiles(layout.File, "handlers/me/editme/my_page_edit.gohtml", "handlers/me/editme/avatars.gohtml"))

func MyPageEdit(w http.ResponseWriter, r *http.Request) {
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

	render.Execute(w, myPageEditTemplate, d)
}

func Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := users.FromContext(ctx)
	username := r.FormValue("username")
	bio := r.FormValue("bio")

	_, err := db.DB.Exec("update users set username=?, bio=? where id=?", username, bio, user.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/me", http.StatusSeeOther)
}
