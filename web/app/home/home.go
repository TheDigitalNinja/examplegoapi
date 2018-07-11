package home

import (
	"html/template"
	"net/http"
)

// Handler index.html template
func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../../template/index.html"))
	tmpl.Execute(w, nil)
}
