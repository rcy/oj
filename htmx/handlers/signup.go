package handlers

import (
	"html/template"
	"net/http"
)

var t = template.Must(template.ParseFiles("templates/signup.html"))

func GetSignup(w http.ResponseWriter, r *http.Request) {
	err := t.Execute(w, nil)
	if err != nil {
		Error(w, err.Error(), 500)
	}
}

func PostSignup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		Error(w, err.Error(), 500)
	}

	username := r.FormValue("username")

	if username == "" {
		http.Redirect(w, r, "/signup", http.StatusFound)
	} else {
		// register the user

		http.SetCookie(w, &http.Cookie{Name: "username", Value: username})
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
