package tools

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"oj/db"
	"oj/element/gradient"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/users"
	"strconv"
)

var t = template.Must(template.ParseFiles(layout.File, "handlers/tools/tools_index.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	l, err := layout.GetData(r)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	render.Execute(w, t, struct {
		Layout   layout.Data
		Gradient gradient.Gradient
	}{
		Layout:   l,
		Gradient: l.BackgroundGradient,
	})
}

func Picker(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		render.Error(w, err.Error(), 500)
	}

	g, err := gradientFromUrlValues(r.PostForm)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	t.ExecuteTemplate(w, "picker", struct {
		Gradient gradient.Gradient
	}{
		Gradient: g,
	})
}

func SetBackground(w http.ResponseWriter, r *http.Request) {
	user := users.Current(r)

	err := r.ParseForm()
	if err != nil {
		render.Error(w, err.Error(), 500)
	}

	g, err := gradientFromUrlValues(r.PostForm)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	encodedGradient, _ := json.Marshal(g)
	_, err = db.DB.Exec("insert into gradients(user_id, gradient) values(?, ?)", user.ID, encodedGradient)
	if err != nil {
		render.Error(w, err.Error(), 500)
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
