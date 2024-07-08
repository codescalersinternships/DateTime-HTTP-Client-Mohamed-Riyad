package pkg

import (
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Url  string
	Port string
}
type Client struct {
	config Config
	client *http.Client
}

func NewConfig(url string, port string) *Config {
	return &Config{url, port}
}
func NewClient(config Config) *Client {
	return &Client{
		config: config,
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}

}
func (c *Client) GetDateTime(endPoint string) (*http.Response, error) {
	return retry(10, 100*time.Millisecond, func() (*http.Response, error) {
		return c.request(endPoint)
	})
}

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
