package fun

import (
	_ "embed"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
)

//go:embed page.gohtml
var pageContent string

var MyPageTemplate = layout.MustParse(pageContent)

func Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l, err := layout.FromContext(ctx)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	d := struct {
		Layout layout.Data
	}{
		Layout: l,
	}

	render.Execute(w, MyPageTemplate, d)
}
