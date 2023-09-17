package templatehelpers

import (
	"html/template"
	"time"

	"github.com/hako/durafmt"
)

var FuncMap = template.FuncMap{
	"fromNow": func(t time.Time) string {
		return durafmt.Parse(time.Now().Sub(t)).LimitFirstN(1).String()
	},
	"odd": func(i, j int) int {
		return (i + j) % 2
	},
}
