package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HomeHandler TODO
func HomeHandler(w http.ResponseWriter, r *http.Request) {}

// HealthCheckHandler respond as up
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"STATUS": "UP"}`)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/Status", HealthCheckHandler)
	log.Print("Staring server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
