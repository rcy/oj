package parent

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"oj/handlers"
	"oj/models/users"

	"github.com/go-chi/chi/v5"
)

func Route(r chi.Router) {
	r.Get("/", index)

	r.Post("/kids", createKid)
}

var t = template.Must(template.ParseFiles("handlers/layout.html", "handlers/parent/index.html"))

func index(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)
	kids, _ := user.Kids()
	err := t.Execute(w, struct {
		User users.User
		Kids []users.User
	}{
		User: user,
		Kids: kids,
	})
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}
}

func createKid(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	username := r.PostForm.Get("username")

	kid, err := users.FindByUsername(username)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, err.Error(), 500)
		return
	}
	if kid != nil {
		http.Error(w, "username taken", http.StatusConflict)
		return
	}

	kid, err = user.CreateKid(username)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Printf("kid: %v", kid)
	http.Redirect(w, r, "/parent", http.StatusSeeOther)
}
