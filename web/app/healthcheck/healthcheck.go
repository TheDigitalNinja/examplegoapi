package healthcheck

import (
	"io"
	"net/http"
)

// Handler respond as up
func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"STATUS": "UP"}`)
}
