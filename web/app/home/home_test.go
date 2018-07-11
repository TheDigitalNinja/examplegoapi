package home

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHomeHandler /
func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handler)
	handler.ServeHTTP(rr, req)

	// Should be StatusOk
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Should be type text/html
	if ctype := rr.Header().Get("Content-Type"); ctype != "text/html; charset=utf-8" {
		t.Errorf("content type header does not match: got %v want %v", ctype, "text/html; charset=utf-8")
	}

}
