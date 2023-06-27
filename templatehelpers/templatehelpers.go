package templatehelpers

import (
	"html/template"
	"time"

	"github.com/hako/durafmt"
)

var FuncMap = template.FuncMap{
	"fromNow": func(t time.Time) string {
		return durafmt.Parse(time.Now().Sub(t)).LimitFirstN(2).String()
	},
}
