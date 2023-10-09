package stickers

import (
	_ "embed"
	"net/http"
	"net/url"
	"oj/db"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/models/users"
	"time"

	goduckgo "github.com/minoplhy/duckduckgo-images-api"
)

var (
	//go:embed page.gohtml
	pageContent  string
	pageTemplate = layout.MustParse(pageContent)
)

type Image struct {
	ID        int64
	CreatedAt time.Time `db:"created_at"`
	UserID    int64     `db:"user_id"`
	URL       string    `db:"url"`
}

func Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l, err := layout.FromContext(ctx)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var images []Image
	err = db.DB.Select(&images, `select * from images where user_id = ? order by created_at desc`, l.User.ID)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d := struct {
		Layout layout.Data
		Images []Image
	}{
		Layout: l,
		Images: images,
	}

	render.Execute(w, pageTemplate, d)
}

func Submit(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")

	keyword := url.QueryEscape("cartoon " + query)

	result := goduckgo.Search(goduckgo.Query{Keyword: keyword})

	render.ExecuteNamed(w, pageTemplate, "result", struct {
		URL       string
		Thumbnail string
	}{
		URL:       result.Results[0].Image,
		Thumbnail: result.Results[0].Thumbnail,
	})
}

func SaveSticker(w http.ResponseWriter, r *http.Request) {
	user := users.FromContext(r.Context())

	url := r.FormValue("url")

	img := Image{URL: url}

	_, err := db.DB.Exec(`insert into images(url, user_id) values(?,?)`, url, user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.ExecuteNamed(w, pageTemplate, "saveSticker", img)
}
