package render

import (
	"html/template"
	"log"
	"net/http"
)

var t = template.Must(template.ParseFiles("handlers/error.html"))

func Error(w http.ResponseWriter, msg string, code int) {
	log.Printf("%d: %s", code, msg)
	w.WriteHeader(code)
	Execute(w, t, struct {
		Message string
		Code    int
	}{
		Message: msg,
		Code:    code,
	})
}
