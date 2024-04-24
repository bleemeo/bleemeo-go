package bleemeo

import (
	"context"
	"encoding/json"
)

type Client struct{}

func NewClient(username, password string) *Client {
	return new(Client)
}

// Get the resource with the given id, with only the given fields, if not nil.
func (c *Client) Get(ctx context.Context, resource, id string, fields []string) (json.RawMessage, error) {
	return nil, nil
}

// List the resources that match given params at the given page, with the given page size.
func (c *Client) List(ctx context.Context, resource string, page, pageSize int, params Params) (ResultsPage, error) {
	return ResultsPage{}, nil
}

// Iterate over resources that match given params.
func (c *Client) Iterator(ctx context.Context, resource string, params Params) Iterator {
	return nil
}

// Create a resource with the given body, and returns only the given fields, if not nil.
func (c *Client) Create(ctx context.Context, resource string, body Body, fields []string) (json.RawMessage, error) {
	return nil, nil
}

// Update the resource with the given id, with the given body, and returns only the given fields, if not nil.
func (c *Client) Update(ctx context.Context, resource, id string, body Body, fields []string) (json.RawMessage, error) {
	return nil, nil
}

// Delete the resource with the given id.
func (c *Client) Delete(ctx context.Context, resource, id string) error {
	return nil
}
