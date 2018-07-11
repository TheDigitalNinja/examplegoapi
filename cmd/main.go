package main

import (
	"examplegoapi/web/app/healthcheck"

	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HomeHandler TODO
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../web/template/index.html"))
	tmpl.Execute(w, nil)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/Status", healthcheck.Handler)
	log.Print("Staring server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
