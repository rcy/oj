package tools

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"oj/db"
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

	g, err := gradientFromUrlValues(r.PostForm)
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
	user := users.Current(r)

	err := r.ParseForm()
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}

	g, err := gradientFromUrlValues(r.PostForm)
	if err != nil {
		handlers.Error(w, err.Error(), 500)
		return
	}

	encodedGradient, _ := json.Marshal(g)
	_, err = db.DB.Exec("insert into gradients(user_id, gradient) values(?, ?)", user.ID, encodedGradient)
	if err != nil {
		handlers.Error(w, err.Error(), 500)
		return
	}

	style := fmt.Sprintf("body { background: %s; }", g.Render())

	w.Write([]byte(style))
}

// Return a Gradient from a parsed form
func gradientFromUrlValues(f url.Values) (gradient.Gradient, error) {
	gradientType := f.Get("gradientType")
	repeat := f.Get("repeat") == "on"
	colors := f["color"]

	// convert []string to []int
	positions := make([]int, len(f["percent"]))
	for i, p := range f["percent"] {
		positions[i], _ = strconv.Atoi(p)
	}

	if len(colors) != len(positions) {
		return gradient.Gradient{}, fmt.Errorf("colors/positions length mismatch")
	}

	degrees, err := strconv.Atoi(f.Get("degrees"))
	if err != nil {
		return gradient.Gradient{}, err
	}
	return gradient.Gradient{
		Type:      gradientType,
		Repeat:    repeat,
		Degrees:   degrees,
		Colors:    colors,
		Positions: positions,
	}, nil
}
