package handlers

import (
	"html/template"
	"net/http"
	"oj/models/users"
)

var homeTemplate = template.Must(template.ParseFiles("handlers/layout.html", "handlers/home.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	user, err := users.FindById(r.Context().Value("userID").(int64))
	err = homeTemplate.Execute(w, struct{ Username string }{Username: user.Username})
	if err != nil {
		Error(w, err.Error(), 500)
	}
}
