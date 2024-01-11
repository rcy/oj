package md

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
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

func Markdown(text string) template.HTML {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(text), &buf); err != nil {
		panic(err)
	}

	return template.HTML(buf.Bytes())
}
