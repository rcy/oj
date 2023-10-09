package header

import (
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
)

var t = layout.MustParse()

func Header(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l, err := layout.FromContext(ctx)
	if err != nil {
		render.Error(w, err.Error(), 500)
		return
	}

	render.ExecuteNamed(w, t, "header", struct{ Layout layout.Data }{l})
}
