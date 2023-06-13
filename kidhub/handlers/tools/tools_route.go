package tools

import (
	"fmt"
	"html/template"
	"net/http"
	"oj/handlers"
	"oj/models/users"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func Route(r chi.Router) {
	r.Get("/", index)
	r.Post("/picker", picker)
}

var t = template.Must(template.ParseFiles("handlers/layout.html", "handlers/tools/tools_index.html"))

type Stop struct {
	Color   string
	Percent int
}

func index(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)
	stops := []Stop{
		{Color: "#ffffff", Percent: 0},
		// {Color: "#000000", Percent: 33},
		// {Color: "#00ff00", Percent: 66},
		{Color: "#0000ff", Percent: 100},
	}
	degrees := 90
	repeat := false

	err := t.Execute(w, struct {
		User            users.User
		GradientType    string
		Repeat          bool
		Stops           []Stop
		Degrees         int
		GradientBar     template.CSS
		GradientPreview template.CSS
	}{
		User:            user,
		GradientType:    "linear",
		Repeat:          repeat,
		Stops:           stops,
		Degrees:         degrees,
		GradientBar:     gradientFromStops("linear", false, 90, stops),
		GradientPreview: gradientFromStops("linear", repeat, degrees, stops),
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

	stops := []Stop{}

	// zip colors and percents into stops
	for i, c := range colors {
		p, _ := strconv.Atoi(percents[i])
		stops = append(stops, Stop{Color: c, Percent: p})
	}

	t.ExecuteTemplate(w, "picker", struct {
		GradientType    string
		Repeat          bool
		Stops           []Stop
		Degrees         int
		GradientBar     template.CSS
		GradientPreview template.CSS
	}{
		GradientType:    gradientType,
		Repeat:          repeat,
		Stops:           stops,
		Degrees:         degrees,
		GradientBar:     gradientFromStops("linear", false, 90, stops),
		GradientPreview: gradientFromStops(gradientType, repeat, degrees, stops),
	})
}

func gradientFromStops(gradientType string, repeating bool, deg int, stops []Stop) template.CSS {
	var params []string

	if repeating {
		gradientType = "repeating-" + gradientType
	}

	switch gradientType {
	case "linear":
		for _, s := range stops {
			params = append(params, fmt.Sprintf("%s %d%%", s.Color, s.Percent))
		}
		return template.CSS(fmt.Sprintf(`linear-gradient(%ddeg, %s)`, deg, strings.Join(params, ",")))

	case "radial":
		for _, s := range stops {
			params = append(params, fmt.Sprintf("%s %d%%", s.Color, s.Percent))
		}
		return template.CSS(fmt.Sprintf(`radial-gradient(%s)`, strings.Join(params, ",")))

	case "conic":
		for _, s := range stops {
			params = append(params, fmt.Sprintf("%s %d%%", s.Color, s.Percent))
		}
		return template.CSS(fmt.Sprintf(`conic-gradient(from %ddeg, %s)`, deg, strings.Join(params, ",")))

	case "repeating-linear":
		for _, s := range stops {
			params = append(params, fmt.Sprintf("%s %dpx", s.Color, s.Percent))
		}
		return template.CSS(fmt.Sprintf(`repeating-linear-gradient(%ddeg, %s)`, deg, strings.Join(params, ",")))
	case "repeating-radial":
		for _, s := range stops {
			params = append(params, fmt.Sprintf("%s %dpx", s.Color, s.Percent))
		}
		return template.CSS(fmt.Sprintf(`repeating-radial-gradient(%s)`, strings.Join(params, ",")))
	case "repeating-conic":
		for _, s := range stops {
			params = append(params, fmt.Sprintf("%s %d%%", s.Color, s.Percent/4))
		}
		return template.CSS(fmt.Sprintf(`repeating-conic-gradient(from %ddeg, %s)`, deg, strings.Join(params, ",")))
	default:
		return template.CSS("black")
	}
}
