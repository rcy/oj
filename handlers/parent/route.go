package parent

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/users"
)

var t = template.Must(template.ParseFiles(layout.File, "handlers/parent/index.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	l, err := layout.GetData(r)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	kids, _ := l.User.Kids()
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	err = t.Execute(w, struct {
		Layout layout.Data
		User   users.User
		Kids   []users.User
	}{
		Layout: l,
		User:   l.User,
		Kids:   kids,
	})
	if err != nil {
		render.Error(w, err.Error(), 500)
	}
}

func CreateKid(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)
	err := r.ParseForm()
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}
	username := r.PostForm.Get("username")

	kid, err := users.FindByUsername(username)
	if err != nil && err != sql.ErrNoRows {
		render.Error(w, err.Error(), 500)
		return
	}
	if kid != nil {
		render.Error(w, "username taken", http.StatusConflict)
		return
	}

	kid, err = user.CreateKid(username)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	log.Printf("kid: %v", kid)
	http.Redirect(w, r, "/parent", http.StatusSeeOther)
}
