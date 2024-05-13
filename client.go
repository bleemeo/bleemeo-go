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
	defaultEndpoint      = "https://api.bleemeo.com"
	defaultOAuthClientID = "1fc6de3e-8750-472e-baea-3ba22bb4eb56"
	defaultUserAgent     = "Bleemeo Go Client"
)

// Client is a helper to interact with the Bleemeo API,
// providing methods to retrieve, list, create, update and delete resources.
type Client struct {
	username, password  string
	endpoint            string
	oAuthClientID       string
	oAuthClientSecret   string
	oAuthInitialRefresh string
	client              *http.Client
	headers             map[string]string

	epURL        *url.URL
	authProvider *authenticationProvider
}

// NewClient initializes a Bleemeo API client with the given options.
// The credentials must be provided by some option.
// The option WithConfigurationFromEnv() might be useful for a default configuration.
func NewClient(opts ...ClientOption) (*Client, error) {
	c := &Client{
		endpoint:      defaultEndpoint,
		oAuthClientID: defaultOAuthClientID,
		client:        new(http.Client),
		headers:       map[string]string{"User-Agent": defaultUserAgent},
	}

	for _, opt := range opts {
		opt(c)
	}

	epURL, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid endpoint URL: %w", err)
	}

	c.epURL = epURL

	if c.oAuthInitialRefresh != "" {
		c.authProvider =
			newRefreshAuthProvider(c.endpoint, c.oAuthClientID, c.oAuthClientSecret, c.oAuthInitialRefresh, c.client)
	} else {
		c.authProvider =
			newCredentialsAuthProvider(c.endpoint, c.username, c.password, c.oAuthClientID, c.oAuthClientSecret, c.client)
	}

	return c, nil
}

// Logout revokes the OAuth token, preventing it from being reused.
func (c *Client) Logout(ctx context.Context) error {
	return c.authProvider.logout(ctx, c.endpoint)
}

// Get the resource with the given id, with only the given fields, if not nil.
func (c *Client) Get(ctx context.Context, resource Resource, id string, fields Fields) (json.RawMessage, error) {
	reqURI := fmt.Sprintf("%s/%s/", resource, id)
	params := Params{"fields": strings.Join(fields, ",")}

	return unmarshalResponse(c.Do(ctx, http.MethodGet, reqURI, params, true, nil))
}

// GetPage returns a list of resources that match given params at the given page, with the given page size.
// To collect all resources matching params (i.e., instead of querying all pages),
// prefer using Iterator() which is faster.
func (c *Client) GetPage(
	ctx context.Context, resource Resource, page, pageSize int, params Params,
) (ResultsPage, error) {
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
		return ResultsPage{}, &JSONUnmarshalError{
			jsonError: &jsonError{
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
func (c *Client) Do(
	ctx context.Context, method, reqURI string, params Params, authenticated bool, body io.Reader,
) (int, []byte, error) {
	reqURL, err := c.epURL.Parse(reqURI)
	if err != nil {
		return 0, nil, fmt.Errorf("bad request URI: %w", err)
	}

	q := reqURL.Query()

	for key, value := range params {
		q.Set(key, value)
	}

	reqURL.RawQuery = q.Encode()

	resp, err := c.do(ctx, method, reqURL.String(), authenticated, body)
	if err != nil {
		return 0, nil, fmt.Errorf("request execution failed: %w", err)
	}

	if resp.StatusCode == 401 {
		cleanupResponse(resp)

		err = c.authProvider.refetchToken(ctx)
		if err != nil {
			return 0, nil, fmt.Errorf("failed to refetch token: %w", err)
		}

		resp, err = c.do(ctx, method, reqURL.String(), authenticated, body) //nolint:bodyclose
		if err != nil {
			return 0, nil, fmt.Errorf("request execution retry failed: %w", err)
		}
	}

	defer cleanupResponse(resp)

	if resp.StatusCode >= 500 {
		return resp.StatusCode, nil, &ServerError{
			apiError: &apiError{
				ReqPath:     reqURL.Path,
				StatusCode:  resp.StatusCode,
				ContentType: resp.Header.Get("Content-Type"),
				Message:     resp.Status,
				Response:    readBodyStart(resp.Body),
			},
		}
	}

	if resp.StatusCode >= 400 {
		bodyStart := readBodyStart(resp.Body)
		clientErr := ClientError{
			apiError: &apiError{
				ReqPath:     reqURL.Path,
				StatusCode:  resp.StatusCode,
				ContentType: resp.Header.Get("Content-Type"),
				Message:     resp.Status,
				Response:    bodyStart,
			},
		}

		if resp.StatusCode == 401 {
			var respBody struct {
				Detail   string `json:"detail"`
				Code     string `json:"code"`
				Messages []struct {
					TokenClass string `json:"token_class"`
					TokenType  string `json:"token_type"`
					Message    string `json:"message"`
				} `json:"messages"`
			}

			err = json.Unmarshal(bodyStart, &respBody)
			if err != nil {
				clientErr.Err = &JSONUnmarshalError{
					&jsonError{
						Err:      err,
						DataKind: JsonErrorDataKind_401Details,
						Data:     bodyStart,
					},
				}
			} else {
				if len(respBody.Messages) > 0 {
					clientErr.Message = respBody.Messages[0].Message
				} else {
					clientErr.Message = respBody.Detail // probably less explicit than the above message
				}

				return resp.StatusCode, nil, &AuthError{
					ClientError: &clientErr,
					ErrorCode:   respBody.Code,
				}
			}
		} else if resp.StatusCode == 404 {
			clientErr.Err = fmt.Errorf("%w: %s", ErrResourceNotFound, reqURL.Path)
		}

		return resp.StatusCode, nil, &clientErr
	}

	respBuf := new(bytes.Buffer)
	if resp.ContentLength > 0 {
		respBuf.Grow(int(resp.ContentLength))
	}

	_, err = respBuf.ReadFrom(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, &ServerError{
			apiError: &apiError{
				ReqPath:     reqURL.Path,
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

func (c *Client) do(
	ctx context.Context, method, reqURL string, authenticated bool, reqBody io.Reader,
) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, reqURL, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to parse request: %w", err)
	}

	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if authenticated {
		err = c.authProvider.injectHeader(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	for header, value := range c.headers {
		req.Header.Set(header, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return resp, nil
}
