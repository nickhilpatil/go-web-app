package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"os"
	"log"
	"io/ioutil"
)

// Mocking http.ServeFile for testing purpose
func TestMain(t *testing.T) {
	// Temporary override of http.ServeFile to return mocked HTML content
	originalServeFile := http.ServeFile
	http.ServeFile = func(w http.ResponseWriter, r *http.Request, filename string) error {
		if filename == "static/home.html" {
			// Serve mocked content directly without accessing the file system
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write([]byte("<html><body>Mocked Home Page</body></html>"))
			return nil
		}
		// For any other file request, we can return an error or handle it
		return os.ErrNotExist
	}
	// Remember to restore the original ServeFile function after the test
	defer func() { http.ServeFile = originalServeFile }() // Restore after test is complete

	// Create request to the /home route
	req, err := http.NewRequest("GET", "/home", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Use the handler for the /home route
	handler := http.HandlerFunc(homePage)

	// Serve the request
	handler.ServeHTTP(rr, req)

	// Check if status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify the content-type header
	expected := "text/html; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("handler returned unexpected content type: got %v want %v", contentType, expected)
	}

	// Verify that the mocked content is returned
	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	expectedBody := "<html><body>Mocked Home Page</body></html>"
	if string(body) != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v", string(body), expectedBody)
	}
}
