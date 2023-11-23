package editme

import (
	_ "embed"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/internal/middleware/auth"
)

var (
	//go:embed my_page_edit.gohtml
	pageContent        string
	myPageEditTemplate = layout.MustParse(pageContent, AvatarContent)
)

func MyPageEdit(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())

	d := struct {
		Layout layout.Data
		User   api.User
	}{
		Layout: l,
		User:   l.User,
	}

	render.Execute(w, myPageEditTemplate, d)
}

func Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := auth.FromContext(ctx)
	username := r.FormValue("username")
	bio := r.FormValue("bio")

	_, err := db.DB.Exec("update users set username=?, bio=? where id=?", username, bio, user.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/me", http.StatusSeeOther)
}
