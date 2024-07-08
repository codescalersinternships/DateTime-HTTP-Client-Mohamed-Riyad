package pkg

import (
	"fmt"
	"net/http"
	"time"
)

// Config holds the configuration parameters for the client.
type Config struct {
	Url  string // URL where the API is hosted
	Port string // Port number
}

// Client is an HTTP client that interacts with a remote API.
type Client struct {
	config Config       // Configuration for the client
	client *http.Client // HTTP client with timeout settings
}

// NewConfig creates a new Config object with the provided URL and port.
func NewConfig(url string, port string) *Config {
	return &Config{url, port}
}

// NewClient creates a new Client object with the provided configuration.
func NewClient(config Config) *Client {
	return &Client{
		config: config,
		client: &http.Client{
			Timeout: time.Second * 10, // Timeout set to 10 seconds
		},
	}
}

// GetDateTime performs an HTTP GET request to retrieve data from the specified endpoint.
func (c *Client) GetDateTime(endPoint string) (*http.Response, error) {
	// Retry logic with exponential backoff
	return retry(10, 100*time.Millisecond, func() (*http.Response, error) {
		return c.request(endPoint)
	})
}

// request performs an HTTP GET request to the endpoint using client's configuration.
func (c *Client) request(endPoint string) (*http.Response, error) {
	url := c.config.Url + endPoint
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// retry executes the provided function with retry logic and returns the response or error.
func retry(attempts int, sleep time.Duration, fn func() (*http.Response, error)) (*http.Response, error) {
	for i := 0; i < attempts; i++ {
		resp, err := fn()
		if err == nil {
			return resp, nil
		}
		time.Sleep(sleep)
	}
	return nil, fmt.Errorf("timed out waiting for response")
}
