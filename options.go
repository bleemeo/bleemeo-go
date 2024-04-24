package bleemeo

import (
	"net/http"
	"os"
)

type ClientOption func(*Client)

func WithCredentials(username, password string) ClientOption {
	return func(c *Client) {
		c.username = username
		c.password = password
	}
}

func WithCredentialsFromEnv() ClientOption {
	return func(c *Client) {
		c.username = os.Getenv("BLEEMEO_USER")
		c.password = os.Getenv("BLEEMEO_PASSWORD")

	}
}

func WithEndpoint(endpoint string) ClientOption {
	return func(c *Client) {
		c.endpoint = endpoint
	}
}

func WithOAuthClientID(clientID string) ClientOption {
	return func(c *Client) {
		c.oAuthClientID = clientID
	}
}

func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.client = client
	}
}
