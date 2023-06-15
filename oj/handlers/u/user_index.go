package u

import (
	"html/template"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/users"
)

var uit = template.Must(template.ParseFiles(layout.File, "handlers/u/people.html"))

func UserIndex(w http.ResponseWriter, r *http.Request) {
	l, err := layout.GetData(r)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	allUsers, err := users.FindAll()
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	d := struct {
		Layout   layout.Data
		User     users.User
		AllUsers []users.User
	}{
		Layout:   l,
		User:     l.User,
		AllUsers: allUsers,
	}

	render.Execute(w, uit, d)
}
