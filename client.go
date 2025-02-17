// Copyright 2015-2025 Bleemeo
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
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"golang.org/x/oauth2"
)

const (
	defaultEndpoint      = "https://api.bleemeo.com"
	defaultOAuthClientID = "1fc6de3e-8750-472e-baea-3ba22bb4eb56"
	defaultUserAgent     = "Bleemeo Go Client"
	// If the throttle delay is less than this, automatically retry the requests.
	defaultThrottleMaxAutoRetryDelay = time.Minute
)

// Client is a helper to interact with the Bleemeo API,
// providing methods to retrieve, list, create, update and delete resources.
type Client struct {
	username, password        string
	endpoint                  string
	oAuthClientID             string
	oAuthClientSecret         string
	oAuthInitialRefresh       string
	client                    *http.Client
	newOAuthTokenCallback     func(token *oauth2.Token)
	headers                   map[string]string
	throttleMaxAutoRetryDelay time.Duration

	epURL        *url.URL
	authProvider *authenticationProvider

	l                sync.Mutex
	throttleDeadline time.Time
}

// NewClient initializes a Bleemeo API client with the given options.
// The credentials must be provided by some option.
// The option WithConfigurationFromEnv() might be useful for a default configuration.
//
// See the README (https://github.com/bleemeo/bleemeo-go/#configuration) for all available options.
//
// The Client will take care of obtaining and refreshing the OAuth token
// to authenticate against the Bleemeo API.
func NewClient(opts ...ClientOption) (*Client, error) {
	c := &Client{
		endpoint:                  defaultEndpoint,
		oAuthClientID:             defaultOAuthClientID,
		client:                    new(http.Client),
		headers:                   map[string]string{"User-Agent": defaultUserAgent},
		throttleMaxAutoRetryDelay: defaultThrottleMaxAutoRetryDelay,
	}

	for _, opt := range opts {
		if opt != nil {
			opt(c)
		}
	}

	if c.username == "" && c.oAuthInitialRefresh == "" {
		return nil, ErrNoAuthMeanProvided
	}

	epURL, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid endpoint URL: %w", err)
	}

	c.epURL = epURL
	c.authProvider = newAuthenticationProvider(
		c.epURL, c.username, c.password, c.oAuthInitialRefresh, c.oAuthClientID, c.oAuthClientSecret, c.client,
	)

	if c.newOAuthTokenCallback != nil {
		c.authProvider.newOAuthTokenCallback = c.newOAuthTokenCallback
	}

	return c, nil
}

// ThrottleDeadline return the time request should be retried.
func (c *Client) ThrottleDeadline() time.Time {
	c.l.Lock()
	defer c.l.Unlock()

	return c.throttleDeadline
}

// GetToken returns the current OAuth token used by the client,
// or retrieves a new one if the current is invalid.
func (c *Client) GetToken(ctx context.Context) (*oauth2.Token, error) {
	return c.authProvider.Token(ctx)
}

// Logout revokes the OAuth token, preventing it from being reused.
func (c *Client) Logout(ctx context.Context) error {
	return c.authProvider.logout(ctx, c.endpoint)
}

// Get the resource with the given id, with only the given fields, if not nil.
func (c *Client) Get(ctx context.Context, resource Resource, id string, fields ...string) (json.RawMessage, error) {
	reqURI, err := url.JoinPath(resource, id, "/")
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return unmarshalResponse(c.Do(ctx, http.MethodGet, reqURI, paramsFromFields(fields), true, nil))
}

// GetPage returns a list of resources that match given params at the given page,
// as pages of the given size.
// To collect all resources matching params (i.e., instead of querying all pages),
// prefer using Iterator() which is faster.
func (c *Client) GetPage(
	ctx context.Context, resource Resource, page, pageSize int, params url.Values,
) (ResultsPage, error) {
	params = cloneMap(params) // avoid mutation of given params
	params.Set("page", strconv.Itoa(page))
	params.Set("page_size", strconv.Itoa(pageSize))

	_, resp, err := c.Do(ctx, http.MethodGet, resource, params, true, nil)
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

// Count the number of resources of the given kind matching the given parameters.
func (c *Client) Count(ctx context.Context, resource Resource, params url.Values) (int, error) {
	result, err := c.GetPage(ctx, resource, 1, 0, params)
	if err != nil {
		return 0, err
	}

	return result.Count, nil
}

// Iterator returns all resources that match given params.
// The page size is set to 2500 by default, but can be defined by setting `page_size` in params.
func (c *Client) Iterator(resource Resource, params url.Values) Iterator {
	return newIterator(c, resource, params)
}

// Create a resource with the given body, which may be any value
// that could be converted to JSON, possibly a simple map[string]string.
// Fields expected to be returned can be specified as variadic parameters.
func (c *Client) Create(ctx context.Context, resource Resource, body any, fields ...string) (json.RawMessage, error) {
	bodyReader, err := JSONReaderFrom(body)
	if err != nil {
		return nil, err
	}

	return unmarshalResponse(c.Do(ctx, http.MethodPost, resource, paramsFromFields(fields), true, bodyReader))
}

// Update the resource with the given id, with the given body, which may be any value
// that could be converted to JSON, possibly a simple map[string]string.
// Since the request is sent as a PATCH, only the fields specified in the body will be updated.
// Fields expected to be returned can be specified as variadic parameters.
func (c *Client) Update(
	ctx context.Context, resource Resource, id string, body any, fields ...string,
) (json.RawMessage, error) {
	bodyReader, err := JSONReaderFrom(body)
	if err != nil {
		return nil, err
	}

	reqURI, err := url.JoinPath(resource, id, "/")
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return unmarshalResponse(c.Do(ctx, http.MethodPatch, reqURI, paramsFromFields(fields), true, bodyReader))
}

// Delete the resource with the given id.
func (c *Client) Delete(ctx context.Context, resource Resource, id string) error {
	reqURI, err := url.JoinPath(resource, id, "/")
	if err != nil {
		return err //nolint: wrapcheck
	}

	_, _, err = c.Do(ctx, http.MethodDelete, reqURI, nil, true, nil)

	return err
}

// Do is a lower-level method to build and execute the request according to the given parameters.
// It returns the response status code and body content, or any error that occurred.
//
// When possible, prefer the higher-level Get, GetPage, Iterator, Create, Update and Delete.
func (c *Client) Do(
	ctx context.Context, method, reqURI string, params url.Values, authenticated bool, body io.Reader,
) (int, []byte, error) {
	if delay := time.Until(c.throttleDeadline); delay > 0 {
		return 0, nil, &ThrottleError{
			APIError: &APIError{
				ReqPath:    reqURI,
				StatusCode: http.StatusTooManyRequests,
				Message:    fmt.Sprintf("Too many requests, need to wait for %s", delay),
			},
			Delay: delay,
		}
	}

	req, err := c.ParseRequest(method, reqURI, nil, params, body)
	if err != nil {
		return 0, nil, err
	}

	statusCode, respBody, err := c.doWithErrorHandling(ctx, req, authenticated)
	if throttleErr := new(ThrottleError); errors.As(err, &throttleErr) {
		if throttleErr.Delay <= c.throttleMaxAutoRetryDelay {
			select {
			case <-time.After(throttleErr.Delay):
			case <-ctx.Done():
				return statusCode, nil, ctx.Err() //nolint:wrapcheck
			}

			statusCode, respBody, err = c.doWithErrorHandling(ctx, req.Clone(ctx), authenticated)
		}
	}

	return statusCode, respBody, err
}

// DoRequest sends the given request and returns the response or any error.
// If authenticated is true, the request will be sent with an Authorization header.
// If the API returns a 401 status code, a new token will be fetched and the request will be sent once again.
// It is up to the caller to close the response body.
func (c *Client) DoRequest(ctx context.Context, req *http.Request, authenticated bool) (*http.Response, error) {
	resp, err := c.do(ctx, req, authenticated)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized && authenticated {
		cleanupResponse(resp)

		err = c.authProvider.refetchToken(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to refetch token: %w", err)
		}

		resp, err = c.do(ctx, req.Clone(ctx), authenticated)
		if err != nil {
			return nil, fmt.Errorf("request execution retry failed: %w", err)
		}
	}

	return resp, nil
}

// ParseRequest returns a new [*http.Request] according to the given values.
// The URL may or may not contain a host, if not, the client's endpoint will be used as such.
func (c *Client) ParseRequest(
	method, url string, headers http.Header, params url.Values, body io.Reader,
) (*http.Request, error) {
	reqURL, err := c.epURL.Parse(url)
	if err != nil {
		return nil, fmt.Errorf("bad request URL: %w", err)
	}

	q := reqURL.Query()

	for key, values := range params {
		for _, value := range values {
			q.Add(key, value)
		}
	}

	reqURL.RawQuery = q.Encode()

	req, err := http.NewRequest(method, reqURL.String(), body) //nolint: lll,noctx // The context will be set by the request executor
	if err != nil {
		return nil, fmt.Errorf("can't parse request: %w", err)
	}

	if body != nil && headers.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	for header, value := range c.headers {
		req.Header.Set(header, value)
	}

	for header, values := range headers {
		for _, value := range values {
			req.Header.Add(header, value)
		}
	}

	return req, nil
}

func (c *Client) doWithErrorHandling(ctx context.Context, req *http.Request, authenticated bool) (int, []byte, error) {
	resp, err := c.DoRequest(ctx, req, authenticated)
	if err != nil {
		return 0, nil, fmt.Errorf("request execution failed: %w", err)
	}

	defer cleanupResponse(resp)

	if resp.StatusCode >= 500 {
		return resp.StatusCode, nil, &APIError{
			ReqPath:     req.URL.Path,
			StatusCode:  resp.StatusCode,
			ContentType: resp.Header.Get("Content-Type"),
			Message:     resp.Status,
			Response:    readBodyStart(resp.Body),
		}
	}

	if resp.StatusCode >= 400 {
		bodyStart := readBodyStart(resp.Body)
		apiErr := APIError{
			ReqPath:     req.URL.Path,
			StatusCode:  resp.StatusCode,
			ContentType: resp.Header.Get("Content-Type"),
			Message:     resp.Status,
			Response:    bodyStart,
		}

		switch resp.StatusCode {
		case http.StatusBadRequest:
			var respBody map[string][]string

			err = json.Unmarshal(bodyStart, &respBody)
			if err != nil {
				apiErr.Err = &JSONUnmarshalError{
					&jsonError{
						Err:      err,
						DataKind: JsonErrorDataKind_400Details,
						Data:     bodyStart,
					},
				}
			} else {
				apiErr.Message = "Bad request:" + makeBadRequestMessage(respBody)
			}
		case http.StatusUnauthorized:
			return resp.StatusCode, nil, buildAuthErrorFromBody(&apiErr)
		case http.StatusNotFound:
			apiErr.Err = fmt.Errorf("%w: %s", ErrResourceNotFound, req.URL.Path)
		case http.StatusTooManyRequests:
			delay := 30 * time.Second

			delaySecond, err := strconv.Atoi(resp.Header.Get("Retry-After"))
			if err == nil {
				delay = time.Duration(delaySecond) * time.Second
			}

			c.l.Lock()
			c.throttleDeadline = time.Now().Add(delay)
			c.l.Unlock()

			apiErr.Message = fmt.Sprintf("Too many requests, need to wait for %s", delay)

			return resp.StatusCode, nil, &ThrottleError{
				APIError: &apiErr,
				Delay:    delay,
			}
		}

		return resp.StatusCode, nil, &apiErr
	}

	respBuf := new(bytes.Buffer)
	if resp.ContentLength > 0 {
		respBuf.Grow(int(resp.ContentLength))
	}

	_, err = respBuf.ReadFrom(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, &APIError{
			ReqPath:     req.URL.Path,
			StatusCode:  resp.StatusCode,
			ContentType: resp.Header.Get("Content-Type"),
			Message:     "failed to read response body",
			Err:         err,
			Response:    nil,
		}
	}

	return resp.StatusCode, respBuf.Bytes(), nil
}

func (c *Client) do(ctx context.Context, req *http.Request, authenticated bool) (*http.Response, error) {
	if authenticated {
		err := c.authProvider.injectHeader(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return resp, nil
}
