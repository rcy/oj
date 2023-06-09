package handlers

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	err := homeTemplate.Execute(w, struct{ Username string }{Username: r.Context().Value("username").(string)})
	if err != nil {
		Error(w, err.Error(), 500)
	}
}
