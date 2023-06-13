package tools

import (
	"html/template"
	"net/http"
	"oj/element/gradient"
	"oj/handlers"
	"oj/models/users"
	"strconv"

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

	g := gradient.Gradient{
		Type:      "linear",
		Repeat:    false,
		Degrees:   90,
		Colors:    []string{"#fffff0", "#0fffff"},
		Positions: []string{"0", "100"},
	}

	err := t.Execute(w, struct {
		User     users.User
		Gradient gradient.Gradient
	}{
		User:     user,
		Gradient: g,
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

	repeat := r.PostForm.Get("repeat") == "on"
	gradientType := r.PostForm.Get("gradientType")
	colors := r.PostForm["color"]
	positions := r.PostForm["percent"]
	degrees, _ := strconv.Atoi(r.PostForm.Get("degrees"))

	t.ExecuteTemplate(w, "picker", struct {
		Gradient gradient.Gradient
	}{
		Gradient: gradient.Gradient{
			Type:      gradientType,
			Repeat:    repeat,
			Degrees:   degrees,
			Colors:    colors,
			Positions: positions,
		},
	})
}

func setBackground(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`body { background: red; }`))
}
