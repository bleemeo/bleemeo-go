package bleemeo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Client struct {
	username, password string
	endpoint           string
	oAuthClientID      string
	client             *http.Client
	customHeaders      map[string]string

	epURL        *url.URL
	authProvider authenticationProvider
}

func NewClient(opts ...ClientOption) (*Client, error) {
	c := &Client{
		endpoint:      "https://api.bleemeo.com",
		client:        new(http.Client),
		customHeaders: make(map[string]string),
	}

	for _, opt := range opts {
		opt(c)
	}

	epURL, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, fmt.Errorf("can't parse endpoint URL: %w", err)
	}

	c.epURL = epURL
	c.authProvider = newAuthProvider(c.endpoint, c.username, c.password, c.oAuthClientID)

	return c, nil
}

// Get the resource with the given id, with only the given fields, if not nil.
func (c *Client) Get(ctx context.Context, resource, id string, fields Fields) (json.RawMessage, error) {
	reqURI := fmt.Sprintf("%s/%s/", resource, id)
	params := Params{"fields": strings.Join(fields, ",")}

	return c.Do(ctx, http.MethodGet, reqURI, params, true, nil)
}

// List the resources that match given params at the given page, with the given page size.
func (c *Client) GetPage(ctx context.Context, resource string, page, pageSize int, params Params) (ResultsPage, error) {
	params = maps.Clone(params) // avoid mutation of given params
	params["page"] = strconv.Itoa(page)
	params["page_size"] = strconv.Itoa(pageSize)

	resp, err := c.Do(ctx, http.MethodGet, resource+"/", params, true, nil)
	if err != nil {
		return ResultsPage{}, err
	}

	var resultPage ResultsPage

	err = json.Unmarshal(resp, &resultPage)
	if err != nil {
		return ResultsPage{}, err
	}

	return resultPage, nil
}

// Iterate over resources that match given params.
func (c *Client) Iterator(resource string, params Params) Iterator {
	return newIterator(c, resource, params)
}

// Create a resource with the given body.
func (c *Client) Create(ctx context.Context, resource string, body Body) (json.RawMessage, error) {
	bodyReader, err := readerFrom(body)
	if err != nil {
		return nil, err
	}

	return c.Do(ctx, http.MethodPost, resource+"/", nil, true, bodyReader)
}

// Update the resource with the given id, with the given body.
func (c *Client) Update(ctx context.Context, resource, id string, body Body) (json.RawMessage, error) {
	bodyReader, err := readerFrom(body)
	if err != nil {
		return nil, err
	}

	return c.Do(ctx, http.MethodPatch, fmt.Sprintf("%s/%s/", resource, id), nil, true, bodyReader)
}

// Delete the resource with the given id.
func (c *Client) Delete(ctx context.Context, resource, id string) error {
	reqURI := fmt.Sprintf("%s/%s/", resource, id)
	_, err := c.Do(ctx, http.MethodDelete, reqURI, nil, true, nil)

	return err
}

func (c *Client) Do(ctx context.Context, method, reqURI string, params Params, authenticated bool, body io.Reader) (json.RawMessage, error) {
	reqURL, err := c.epURL.Parse(reqURI)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL.String(), body)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()

	for key, value := range params {
		q.Set(key, value)
	}

	req.URL.RawQuery = q.Encode()

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if authenticated {
		err = c.authProvider.injectHeader(req)
		if err != nil {
			return nil, err
		}
	}

	for header, value := range c.customHeaders {
		req.Header.Set(header, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	raw := make(json.RawMessage, 0, resp.ContentLength)

	err = json.NewDecoder(resp.Body).Decode(&raw)
	if err != nil {
		return nil, err
	}

	return raw, nil
}
