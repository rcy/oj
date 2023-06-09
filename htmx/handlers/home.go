package handlers

import (
	"html/template"
	"net/http"
)

var homeTemplate = template.Must(template.ParseFiles("templates/layout.html", "templates/index.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	err := homeTemplate.Execute(w, struct{ Username string }{Username: r.Context().Value("username").(string)})
	if err != nil {
		Error(w, err.Error(), 500)
	}
}
