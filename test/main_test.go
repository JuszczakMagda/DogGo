package test

import (
	"DogGo/src/server"
	"github.com/jarcoal/httpmock"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	server.Handler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", resp.StatusCode)
	}
	expected := "<h1>DANTE</h1>"
	if string(body) != expected {
		t.Errorf("Expected body %q, got %q", expected, string(body))
	}
}
func TestRandomDogHandler_MockedAPI(t *testing.T) {
	// Activate httpmock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock the external API call
	mockURL := "https://dog.ceo/api/breeds/image/random"
	httpmock.RegisterResponder("GET", mockURL,
		httpmock.NewStringResponder(200, `{
			"message": "https://images.dog.ceo/breeds/hound-afghan/n02088094_1003.jpg",
			"status": "success"
		}`))

	// Create a request to the /doggo route
	req := httptest.NewRequest(http.MethodGet, "/doggo", nil)
	w := httptest.NewRecorder()

	// Call the actual handler
	server.RandomDogHandler(w, req) // See note below

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	bodyStr := string(body)

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", resp.StatusCode)
	}

	if !strings.Contains(bodyStr, "<img") {
		t.Errorf("Expected <img> tag in response body, got: %s", bodyStr)
	}

	expectedURL := "https://images.dog.ceo/breeds/hound-afghan/n02088094_1003.jpg"
	if !strings.Contains(bodyStr, expectedURL) {
		t.Errorf("Expected image URL %q in response body, got: %s", expectedURL, bodyStr)
	}
}

// Helper to normalize and compare HTML-ish strings
func stringContainsIgnoreWhitespace(haystack, needle string) bool {
	return removeWhitespace(haystack) == removeWhitespace(needle)
}

func removeWhitespace(s string) string {
	out := make([]rune, 0, len(s))
	for _, r := range s {
		if r != ' ' && r != '\n' && r != '\t' {
			out = append(out, r)
		}
	}
	return string(out)
}
