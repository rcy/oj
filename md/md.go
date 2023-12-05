package md

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"mvdan.cc/xurls/v2"
)

func RenderString(text string) template.HTML {
	rx := xurls.Relaxed()
	str := rx.ReplaceAllStringFunc(text, func(link string) string {
		withProtocol := link

		if !strings.HasPrefix(link, "https://") && !strings.HasPrefix(link, "http://") {
			withProtocol = "https://" + link
		}
		return fmt.Sprintf(`<a href="%s" target="_blank">%s</a>`, withProtocol, link)
	})

	bm := bluemonday.NewPolicy()
	bm.AllowStandardURLs()
	bm.AllowAttrs("href", "target").OnElements("a")
	bm.AllowElements("p")

	html := bm.SanitizeBytes([]byte(str))
	return template.HTML(fmt.Sprintf("<p>%s</p>\n", html))
}
