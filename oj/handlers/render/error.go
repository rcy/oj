package render

import (
	"log"
	"net/http"
)

func Error(w http.ResponseWriter, msg string, code int) {
	http.Error(w, msg, code)
	log.Printf("%d %s", code, msg)
}
