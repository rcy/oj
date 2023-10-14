package header

import (
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
)

var t = layout.MustParse()

func Header(w http.ResponseWriter, r *http.Request) {
	l := layout.FromContext(r.Context())

	render.ExecuteNamed(w, t, "header", struct{ Layout layout.Data }{l})
}
