package tools

import (
	"html/template"
	"net/http"
	"oj/element/gradient"
	"oj/handlers"
	"oj/models/users"

	"github.com/go-chi/chi/v5"
)

func Route(r chi.Router) {
	r.Get("/", index)
	r.Post("/picker", picker)
	r.Post("/set-background", setBackground)
}

var t = template.Must(template.ParseFiles("handlers/layout.html", "handlers/tools/tools_index.html"))

func index(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)

	err := t.Execute(w, struct {
		User     users.User
		Gradient gradient.Gradient
	}{
		User:     user,
		Gradient: gradient.Default,
	})
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}
}

func picker(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}

	g, err := gradient.ParseForm(r.PostForm)
	if err != nil {
		handlers.Error(w, err.Error(), 500)
		return
	}

	t.ExecuteTemplate(w, "picker", struct {
		Gradient gradient.Gradient
	}{
		Gradient: g,
	})
}

func setBackground(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`body { background: red; }`))
}
