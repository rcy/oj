package admin

import (
	mw "oj/handlers/middleware"

	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Use(mw.EnsureAdmin)
	r.Get("/", Page)
}
