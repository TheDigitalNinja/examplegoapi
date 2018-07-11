package main

import (
	"examplegoapi/web/app/healthcheck"
	"examplegoapi/web/app/home"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home.Handler)
	router.HandleFunc("/Status", healthcheck.Handler)
	log.Print("Staring server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
