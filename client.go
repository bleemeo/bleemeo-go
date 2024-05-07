// Copyright 2015-2024 Bleemeo
//
// bleemeo.com an infrastructure monitoring solution in the Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bleemeo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	defaultEndpoint  = "https://api.bleemeo.com"
	defaultUserAgent = "Bleemeo Go Client"
)

type Client struct {
	username, password  string
	endpoint            string
	oAuthClientID       string
	oAuthClientSecret   string
	oAuthInitialRefresh string
	client              *http.Client
	customHeaders       map[string]string

	epURL        *url.URL
	authProvider authenticationProvider
}

// NewClient initializes a Bleemeo API client with the given options.
// The credentials should be provided by some option.
// The option WithConfigurationFromEnv() might be useful for a default configuration.
func NewClient(opts ...ClientOption) (*Client, error) {
	c := &Client{
		endpoint:      defaultEndpoint,
		client:        new(http.Client),
		customHeaders: map[string]string{"User-Agent": defaultUserAgent},
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.oAuthClientID == "" { // Common pitfall
		return nil, ErrNoOAuthClientIDProvided
	}

	epURL, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid endpoint URL: %w", err)
	}

	c.epURL = epURL

	if c.oAuthInitialRefresh == "" {
		tk, err := newCredentialsAuthProvider(c.endpoint, c.username, c.password, c.oAuthClientID, c.oAuthClientSecret, c.client).Token()
		if err != nil {
			return nil, fmt.Errorf("can't retrieve OAuth token: %w", err)
		}

		c.oAuthInitialRefresh = tk.RefreshToken
	}

	c.authProvider = newRefreshAuthProvider(c.endpoint, c.oAuthClientID, c.oAuthClientSecret, c.oAuthInitialRefresh, c.client)

	return c, nil
}

// Logout revokes the OAuth token and prevents it from being reused.
func (c *Client) Logout(ctx context.Context) error {
	currentToken, err := c.authProvider.Token()
	if err != nil {
		return fmt.Errorf("couldn't retrieve token: %w", err)
	}

	// Revoking the refresh token will also revoke the related access token
	body := strings.NewReader(fmt.Sprintf("client_id=%s&token_type_hint=refresh_token&token=%s", c.oAuthClientID, currentToken.RefreshToken))
	// Temporarily modifying the content type to override application/json
	previousContentType, hadContentType := c.customHeaders["Content-Type"]
	c.customHeaders["Content-Type"] = "application/x-www-form-urlencoded"

	statusCode, _, err := c.Do(ctx, http.MethodPost, "o/revoke_token/", nil, true, body)

	if hadContentType {
		c.customHeaders["Content-Type"] = previousContentType
	} else {
		delete(c.customHeaders, "Content-Type")
	}

	if err != nil {
		return fmt.Errorf("%w: %w", ErrTokenRevoke, err)
	}

	if statusCode != http.StatusOK {
		return fmt.Errorf("%w: server replyed with status code %d", ErrTokenRevoke, statusCode)
	}

	return nil
}

// Get the resource with the given id, with only the given fields, if not nil.
func (c *Client) Get(ctx context.Context, resource Resource, id string, fields Fields) (json.RawMessage, error) {
	reqURI := fmt.Sprintf("%s/%s/", resource, id)
	params := Params{"fields": strings.Join(fields, ",")}

	return unmarshalResponse(c.Do(ctx, http.MethodGet, reqURI, params, true, nil))
}

// GetPage returns a list of resources that match given params at the given page, with the given page size.
// To collect all resources matching params (i.e., instead of querying all pages), prefer using Iterator() which is faster.
func (c *Client) GetPage(ctx context.Context, resource Resource, page, pageSize int, params Params) (ResultsPage, error) {
	params = cloneMap(params) // avoid mutation of given params
	params["page"] = strconv.Itoa(page)
	params["page_size"] = strconv.Itoa(pageSize)

	_, resp, err := c.Do(ctx, http.MethodGet, string(resource+"/"), params, true, nil)
	if err != nil {
		return ResultsPage{}, err
	}

	var resultPage ResultsPage

	err = json.Unmarshal(resp, &resultPage)
	if err != nil {
		return ResultsPage{}, &JsonUnmarshalError{
			jsonError: jsonError{
				Err:      err,
				DataKind: JsonErrorDataKind_ResultPage,
				Data:     resp,
			},
		}
	}

	return resultPage, nil
}

// Iterator returns all resources that match given params.
func (c *Client) Iterator(resource Resource, params Params) Iterator {
	return newIterator(c, resource, params)
}

// Create a resource with the given body.
func (c *Client) Create(ctx context.Context, resource Resource, body Body) (json.RawMessage, error) {
	bodyReader, err := jsonReaderFrom(body)
	if err != nil {
		return nil, err
	}

	return unmarshalResponse(c.Do(ctx, http.MethodPost, string(resource+"/"), nil, true, bodyReader))
}

// Update the resource with the given id, with the given body.
func (c *Client) Update(ctx context.Context, resource Resource, id string, body Body) (json.RawMessage, error) {
	bodyReader, err := jsonReaderFrom(body)
	if err != nil {
		return nil, err
	}

	return unmarshalResponse(c.Do(ctx, http.MethodPatch, fmt.Sprintf("%s/%s/", resource, id), nil, true, bodyReader))
}

// Delete the resource with the given id.
func (c *Client) Delete(ctx context.Context, resource Resource, id string) error {
	reqURI := fmt.Sprintf("%s/%s/", resource, id)
	_, _, err := c.Do(ctx, http.MethodDelete, reqURI, nil, true, nil)

	return err
}

// Do builds and execute the request according to the given parameters.
// It returns the response status code and body content, or any error that occurred.
func (c *Client) Do(ctx context.Context, method, reqURI string, params Params, authenticated bool, body io.Reader) (int, []byte, error) {
	reqURL, err := c.epURL.Parse(reqURI)
	if err != nil {
		return 0, nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL.String(), body)
	if err != nil {
		return 0, nil, err
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
			return 0, nil, err
		}
	}

	for header, value := range c.customHeaders {
		req.Header.Set(header, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return 0, nil, err
	}

	defer func() {
		// Ensure we read the whole response to avoid "Connection reset by peer" on server
		// and ensure HTTP connection can be reused
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	if resp.StatusCode >= 500 {
		return resp.StatusCode, nil, &ServerError{
			apiError: apiError{
				ReqPath:     req.URL.Path,
				StatusCode:  resp.StatusCode,
				ContentType: resp.Header.Get("Content-Type"),
				Message:     resp.Status,
				Response:    readBodyStart(resp.Body),
			},
		}
	}

	if resp.StatusCode >= 400 {
		var underlyingError error

		if resp.StatusCode == 404 {
			underlyingError = fmt.Errorf("%w: %s", ErrResourceNotFound, req.URL.Path)
		}

		return resp.StatusCode, nil, &ClientError{
			apiError: apiError{
				ReqPath:     req.URL.Path,
				StatusCode:  resp.StatusCode,
				ContentType: resp.Header.Get("Content-Type"),
				Message:     resp.Status,
				Err:         underlyingError,
				Response:    readBodyStart(resp.Body),
			},
		}
	}

	respBuf := new(bytes.Buffer)
	respBuf.Grow(int(resp.ContentLength))

	_, err = respBuf.ReadFrom(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, &ServerError{
			apiError: apiError{
				ReqPath:     req.URL.Path,
				StatusCode:  resp.StatusCode,
				ContentType: resp.Header.Get("Content-Type"),
				Message:     "can't read response body",
				Err:         err,
				Response:    nil,
			},
		}
	}

	return resp.StatusCode, respBuf.Bytes(), nil
}
