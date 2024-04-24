package bleemeo

import (
	"context"
	"encoding/json"
	"os"
)

type Client struct {
	username, password string
	endpoint           string
}

func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		username: os.Getenv("BLEEMEO_USER"),
		password: os.Getenv("BLEEMEO_PASSWORD"),
		endpoint: "api.bleemeo.com/v1",
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Get the resource with the given id, with only the given fields, if not nil.
func (c *Client) Get(ctx context.Context, resource, id string, fields []string) (json.RawMessage, error) {
	return nil, nil
}

// List the resources that match given params at the given page, with the given page size.
func (c *Client) GetPage(ctx context.Context, resource string, page, pageSize int, params Params) (ResultsPage, error) {
	return ResultsPage{}, nil
}

// Iterate over resources that match given params.
func (c *Client) Iterator(resource string, params Params) Iterator {
	return nil
}

// Create a resource with the given body.
func (c *Client) Create(ctx context.Context, resource string, body Body) (json.RawMessage, error) {
	return nil, nil
}

// Update the resource with the given id, with the given body.
func (c *Client) Update(ctx context.Context, resource, id string, body Body) (json.RawMessage, error) {
	return nil, nil
}

// Delete the resource with the given id.
func (c *Client) Delete(ctx context.Context, resource, id string) error {
	return nil
}
