package layout

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"oj/db"
	"oj/element/gradient"
	"oj/models/gradients"
	"oj/models/users"
	"oj/templatehelpers"
)

var (
	//go:embed "layout.gohtml"
	layoutContent string
)

func MustParse(templateContent ...string) *template.Template {
	tpl := template.New("layout").Funcs(templatehelpers.FuncMap)

	for i, content := range append([]string{layoutContent}, templateContent...) {
		var err error
		tpl, err = tpl.Parse(content)
		if err != nil {
			fmt.Println(i, content)
			panic(err)
		}
	}
	return tpl
}

const File = "handlers/layout/layout.gohtml"

type Data struct {
	User               users.User
	URL                url.URL
	PageURL            url.URL
	BackgroundGradient gradient.Gradient
	UnreadCount        int
}

func FromRequest(r *http.Request) (Data, error) {
	user := users.FromContext(r.Context())

	backgroundGradient, err := gradients.UserBackground(user.ID)
	if err != nil {
		return Data{}, err
	}

	var unreadCount int
	err = db.DB.Get(&unreadCount, `select count(*) from deliveries where recipient_id = ? and sent_at is null`, user.ID)
	if err != nil {
		return Data{}, err
	}

	return Data{
		User:               user,
		URL:                *r.URL,
		PageURL:            pageURLFromRequest(r),
		BackgroundGradient: *backgroundGradient,
		UnreadCount:        unreadCount,
	}, nil
}

// Return the URL of the current page, which can be different from the request url when called in the context of an htmx/ajax request
func pageURLFromRequest(r *http.Request) url.URL {
	hxCurrentURL := r.Header.Get("HX-Current-URL")
	if hxCurrentURL == "" {
		return *r.URL
	}
	url, err := url.Parse(hxCurrentURL)
	if err != nil {
		return *r.URL
	}
	return *url
}
