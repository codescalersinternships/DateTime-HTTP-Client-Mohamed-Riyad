package pkg

import (
	"fmt"
	"io/ioutil"
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
func (c *Client) GetDateTime(endPoint string) (string, error) {
	return retry(10, 100*time.Millisecond, func() (string, error) {
		return c.request(endPoint)
	})
}

func (c *Client) request(endPoint string) (string, error) {
	url := c.config.Url + endPoint
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
func retry(attempts int, sleep time.Duration, fn func() (string, error)) (string, error) {
	for i := 0; i < attempts; i++ {
		body, err := fn()
		if err == nil {
			return body, nil
		}
		time.Sleep(sleep)
	}
	return "", fmt.Errorf("timed out waiting for response")
}
