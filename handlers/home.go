package handlers

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/me", http.StatusFound)
}
