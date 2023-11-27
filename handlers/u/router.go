package u

import "github.com/go-chi/chi/v5"

func Router(r chi.Router) {
	r.Get("/", Page)
	r.Get("/postcard", Page)
}
