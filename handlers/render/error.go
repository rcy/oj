package render

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"
)

var (
	//go:embed "error.gohtml"
	tContent string
	t        = template.Must(template.New("").Parse(tContent))
)

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
