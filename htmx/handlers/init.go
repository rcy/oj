package handlers

import (
	"html/template"
	"log"
)

var homeTemplate *template.Template

func init() {
	log.Printf("Init Templates")
	homeTemplate = template.Must(template.ParseFiles("templates/index.html"))
}
