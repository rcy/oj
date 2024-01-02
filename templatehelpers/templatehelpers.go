package templatehelpers

import (
	"html/template"
	"oj/md"
	"time"

	"github.com/hako/durafmt"
	"github.com/rcy/durfmt"
)

var FuncMap = template.FuncMap{
	"fromNow": func(t time.Time) string {
		return durafmt.Parse(time.Now().Sub(t)).LimitFirstN(1).String()
	},
	"odd": func(i, j int) int {
		return (i + j) % 2
	},
	"html": func(str string) template.HTML {
		return md.RenderString(str)
	},
	"ago": func(t time.Time) string {
		dur := time.Now().Sub(t)
		if dur < time.Minute {
			return "just now"
		}
		return durfmt.Format(dur) + " ago"
	},
}
