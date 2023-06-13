package tools

import (
	"html/template"
	"net/http"
	"oj/element"
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
	stops := []element.Stop{
		{Color: "#ffffff", Percent: 0},
		// {Color: "#000000", Percent: 33},
		// {Color: "#00ff00", Percent: 66},
		{Color: "#0000ff", Percent: 100},
	}
	degrees := 90
	repeat := false

	err := t.Execute(w, struct {
		User     users.User
		Gradient element.Gradient
	}{
		User: user,
		Gradient: element.Gradient{
			Type:    "linear",
			Repeat:  repeat,
			Stops:   stops,
			Degrees: degrees,
		},
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
	percents := r.PostForm["percent"]
	degrees, _ := strconv.Atoi(r.PostForm.Get("degrees"))

	stops := []element.Stop{}

	// zip colors and percents into stops
	for i, c := range colors {
		p, _ := strconv.Atoi(percents[i])
		stops = append(stops, element.Stop{Color: c, Percent: p})
	}

	t.ExecuteTemplate(w, "picker", struct {
		Gradient element.Gradient
	}{
		Gradient: element.Gradient{
			Type:    gradientType,
			Repeat:  repeat,
			Stops:   stops,
			Degrees: degrees,
		},
	})
}

func setBackground(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`body { background: red; }`))
}
