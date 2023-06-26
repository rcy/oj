package md

import (
	"html/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/microcosm-cc/bluemonday"
)

func RenderString(text string) template.HTML {
	p := parser.NewWithExtensions(parser.NoExtensions | parser.Autolink)
	r := html.NewRenderer(html.RendererOptions{Flags: html.HrefTargetBlank | html.SkipImages})
	bm := bluemonday.NewPolicy()
	bm.AllowStandardURLs()
	bm.AllowAttrs("href", "target").OnElements("a")
	bm.AllowElements("p")

	bytes := markdown.ToHTML([]byte(text), p, r)

	html := bm.SanitizeBytes(bytes)

	return template.HTML(html)
}
