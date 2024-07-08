package pkg

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestNewClient tests the creation of a new Client
func TestNewClient(t *testing.T) {
	config := Config{Url: "http://localhost", Port: "8080"}
	client := NewClient(config)
	if client.config.Url != "http://localhost" || client.config.Port != "8080" {
		t.Errorf("Expected URL: %s and Port: %s, but got URL: %s and Port: %s", "http://localhost", "8080", client.config.Url, client.config.Port)
	}
}

// TestNewConfig tests the creation of a new config
func TestNewConfig(t *testing.T) {
	config := Config{Url: "http://localhost", Port: "8080"}
	if config.Url != "http://localhost" || config.Port != "8080" {
		t.Errorf("Expected URL: %s and Port: %s, but got URL: %s and Port: %s", "http://localhost", "8080", config.Url, config.Port)
	}
}

// TestGetDateTime tests the GetDateTime method of the Client
func TestGetDateTime(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		response := currentTime
		json.NewEncoder(w).Encode(response)
	}))

	// Use the mock server URL in the client config
	config := Config{Url: server.URL, Port: ""}
	client := NewClient(config)

	// Call the GetDateTime method
	resp, err := client.GetDateTime("/datetime")
	if err != nil {
		t.Errorf("error making request: %v", err)
	}

	var responseTime time.Time
	err = json.NewDecoder(resp.Body).Decode(&responseTime)
	if err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	// Compare the parsed time with the expected value
	got := responseTime.Round(time.Minute)
	want := time.Now().Round(time.Minute)

	if !got.Equal(want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
