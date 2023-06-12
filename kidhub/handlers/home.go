package handlers

import (
	"html/template"
	"net/http"
	"oj/models/users"
)

var homeTemplate = template.Must(template.ParseFiles("handlers/layout.html", "handlers/home.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)

	if user.IsParent() {
		http.Redirect(w, r, "/parent", http.StatusFound)
	}

	err := homeTemplate.Execute(w, struct{ User users.User }{User: user})
	if err != nil {
		Error(w, err.Error(), 500)
	}
}
