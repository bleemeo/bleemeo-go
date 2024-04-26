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

type Client struct {
	username, password string
	endpoint           string
	oAuthClientID      string
	client             *http.Client
	customHeaders      map[string]string

	epURL        *url.URL
	authProvider authenticationProvider
}

// NewClient initializes a Bleemeo API client with the given options.
// If an error is returned, it will be of type [url.Error].
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
		return nil, fmt.Errorf("invalid endpoint URL: %w", err)
	}

	c.epURL = epURL
	c.authProvider = newAuthProvider(c.endpoint, c.username, c.password, c.oAuthClientID)

	return c, nil
}

// Get the resource with the given id, with only the given fields, if not nil.
func (c *Client) Get(ctx context.Context, resource, id string, fields Fields) (json.RawMessage, error) {
	reqURI := fmt.Sprintf("%s/%s/", resource, id)
	params := Params{"fields": strings.Join(fields, ",")}

	return unmarshalResponse(c.Do(ctx, http.MethodGet, reqURI, params, true, nil))
}

// List the resources that match given params at the given page, with the given page size.
func (c *Client) GetPage(ctx context.Context, resource string, page, pageSize int, params Params) (ResultsPage, error) {
	params = cloneMap(params) // avoid mutation of given params
	params["page"] = strconv.Itoa(page)
	params["page_size"] = strconv.Itoa(pageSize)

	resp, err := c.Do(ctx, http.MethodGet, resource+"/", params, true, nil)
	if err != nil {
		return ResultsPage{}, err
	}

	var resultPage ResultsPage

	err = json.Unmarshal(resp, &resultPage)
	if err != nil {
		return ResultsPage{}, &JsonUnmarshalError{
			jsonError: jsonError{
				Err:      err,
				DataKind: "result page",
				Data:     resp,
			},
		}
	}

	return resultPage, nil
}

// Iterate over resources that match given params.
func (c *Client) Iterator(resource string, params Params) Iterator {
	return newIterator(c, resource, params)
}

// Create a resource with the given body.
func (c *Client) Create(ctx context.Context, resource string, body Body) (json.RawMessage, error) {
	bodyReader, err := jsonReaderFrom(body)
	if err != nil {
		return nil, err
	}

	return unmarshalResponse(c.Do(ctx, http.MethodPost, resource+"/", nil, true, bodyReader))
}

// Update the resource with the given id, with the given body.
func (c *Client) Update(ctx context.Context, resource, id string, body Body) (json.RawMessage, error) {
	bodyReader, err := jsonReaderFrom(body)
	if err != nil {
		return nil, err
	}

	return unmarshalResponse(c.Do(ctx, http.MethodPatch, fmt.Sprintf("%s/%s/", resource, id), nil, true, bodyReader))
}

// Delete the resource with the given id.
func (c *Client) Delete(ctx context.Context, resource, id string) error {
	reqURI := fmt.Sprintf("%s/%s/", resource, id)
	_, err := c.Do(ctx, http.MethodDelete, reqURI, nil, true, nil)

	return err
}

// Do builds and execute the request according to the given parameters.
// It returns the response body content, or any error that occurred.
func (c *Client) Do(ctx context.Context, method, reqURI string, params Params, authenticated bool, body io.Reader) ([]byte, error) {
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

	defer func() {
		// Ensure we read the whole response to avoid "Connection reset by peer" on server
		// and ensure HTTP connection can be reused
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	if resp.StatusCode >= 500 {
		return nil, &ServerError{
			apiError: apiError{
				ReqPath:     req.URL.Path,
				StatusCode:  resp.StatusCode,
				ContentType: resp.Header.Get("Content-Type"),
				Message:     resp.Status,
				Response:    readBodyStart(resp.Body),
			},
		}
	}

	if resp.StatusCode == 404 {
		// 404 is JSON with a "detail" key
		var details struct {
			Detail string `json:"detail"`
		}

		err = json.NewDecoder(resp.Body).Decode(&details)
		if err != nil {
			// TODO: this read misses all the data already read by the JSON decoder
			content := readBodyStart(resp.Body)

			return nil, &ClientError{
				apiError: apiError{
					ReqPath:     req.URL.Path,
					StatusCode:  404,
					ContentType: resp.Header.Get("Content-Type"),
					Message:     resp.Status,
					Err: &JsonUnmarshalError{
						jsonError: jsonError{
							DataKind: "404 details",
							Err:      err,
							Data:     content,
						},
					},
					Response: content,
				},
			}
		}

		message := resp.Status
		if details.Detail != "" {
			message = details.Detail
		}

		return nil, &ClientError{
			apiError: apiError{
				ReqPath:     req.URL.Path,
				StatusCode:  404,
				ContentType: resp.Header.Get("Content-Type"),
				Message:     message,
				Err:         fmt.Errorf("%w: %s", ErrResourceNotFound, req.URL.Path),
				Response:    readBodyStart(resp.Body),
			},
		}

	}

	if resp.StatusCode >= 400 {
		return nil, &ClientError{
			apiError: apiError{
				ReqPath:     req.URL.Path,
				StatusCode:  resp.StatusCode,
				ContentType: resp.Header.Get("Content-Type"),
				Message:     resp.Status,
				Response:    readBodyStart(resp.Body),
			},
		}
	}

	respBuf := new(bytes.Buffer)
	respBuf.Grow(int(resp.ContentLength))

	_, err = respBuf.ReadFrom(resp.Body)
	if err != nil {
		return nil, &ServerError{
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

	return respBuf.Bytes(), nil
}
