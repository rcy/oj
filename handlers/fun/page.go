package fun

import (
	_ "embed"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
)

var (
	//go:embed page.gohtml
	pageContent string

	MyPageTemplate = layout.MustParse(pageContent)
)

func Page(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())

	d := struct {
		Layout layout.Data
	}{
		Layout: l,
	}

	render.Execute(w, MyPageTemplate, d)
}
