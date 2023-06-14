package u

import (
	"html/template"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/users"
)

var t = template.Must(template.ParseFiles(layout.File, "handlers/u/user_page.html"))

func UserPage(w http.ResponseWriter, r *http.Request) {
	l, err := layout.GetData(r)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	d := struct {
		Layout layout.Data
		User   users.User
	}{
		Layout: l,
		User:   l.User,
	}

	render.Execute(w, t, d)
}
