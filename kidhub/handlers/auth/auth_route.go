package auth

import (
	"html/template"
	"net/http"
	"oj/handlers"
	"time"

	"github.com/go-chi/chi/v5"
)

func Route(r chi.Router) {
	r.Get("/", welcome)
	r.Get("/kids", welcomeKids)
	r.Get("/parents", welcomeParents)

	r.Get("/signup", getSignup)
	r.Post("/signup", postSignup)
	r.Get("/signout", signout)
}

var welcomeTemplate = template.Must(template.ParseFiles(
	"handlers/auth/layout.html",
	"handlers/auth/welcome.html",
))

func welcome(w http.ResponseWriter, r *http.Request) {
	err := welcomeTemplate.Execute(w, nil)
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}
}

var welcomeKidsTemplate = template.Must(template.ParseFiles(
	"handlers/auth/layout.html",
	"handlers/auth/welcome_kids.html",
))

func welcomeKids(w http.ResponseWriter, r *http.Request) {
	err := welcomeKidsTemplate.Execute(w, nil)
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}
}

var welcomeParentsTemplate = template.Must(template.ParseFiles(
	"handlers/auth/layout.html",
	"handlers/auth/welcome_parents.html",
))

func welcomeParents(w http.ResponseWriter, r *http.Request) {
	err := welcomeParentsTemplate.Execute(w, nil)
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}
}

func signout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "username", Path: "/", Expires: time.Now().Add(-time.Hour)})
	http.Redirect(w, r, "/", http.StatusFound)
}

var t = template.Must(template.ParseFiles("handlers/auth/auth_signup.html"))

func getSignup(w http.ResponseWriter, r *http.Request) {
	err := t.Execute(w, nil)
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}
}

func postSignup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}

	username := r.FormValue("username")

	if username == "" {
		http.Redirect(w, r, "/signup", http.StatusFound)
	} else {
		// register the user

		http.SetCookie(w, &http.Cookie{Name: "username", Path: "/", Value: username})
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
