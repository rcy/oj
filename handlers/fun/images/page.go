package images

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/templatehelpers"

	goduckgo "github.com/minoplhy/duckduckgo-images-api"
)

var pageTemplate = template.Must(template.New("layout.gohtml").Funcs(templatehelpers.FuncMap).ParseFiles(layout.File, "handlers/fun/images/page.gohtml"))

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

	render.Execute(w, pageTemplate, d)
}

func Submit(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")

	result := goduckgo.Search(goduckgo.Query{Keyword: fmt.Sprintf("cartoon %s", query)})
	log.Printf("%v", result)

	render.ExecuteNamed(w, pageTemplate, "result", struct {
		Image     string
		Thumbnail string
	}{
		Image:     result.Results[0].Image,
		Thumbnail: result.Results[0].Thumbnail,
	})
}
