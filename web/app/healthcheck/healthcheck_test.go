package healthcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHealthCheckHandler /Status
func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/Status", nil)
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

	// Should be type json
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("content type header does not match: got %v want %v", ctype, "application/json")
	}

	// Should be {"STATUS": "UP"}
	expected := `{"STATUS": "UP"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
