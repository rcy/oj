package oj

import (
	"net/http"
)

func Redirect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Host == "octopusjr.ca" {
			http.Redirect(w, r, "https://kable.ca", http.StatusMovedPermanently)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
