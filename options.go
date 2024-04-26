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
	"net/http"
	"os"
)

// A ClientOption can be used to customize the [Client].
type ClientOption func(*Client)

// WithCredentials will make the client use the given credentials.
func WithCredentials(username, password string) ClientOption {
	return func(c *Client) {
		c.username = username
		c.password = password
	}
}

// WithCredentialsFromEnv will make the client retrieve and use credentials
// from the "BLEEMEO_USER" and "BLEEMEO_PASSWORD" environment variables.
func WithCredentialsFromEnv() ClientOption {
	return func(c *Client) {
		c.username = os.Getenv("BLEEMEO_USER")
		c.password = os.Getenv("BLEEMEO_PASSWORD")

	}
}

// WithEndpoint will make the client use the given endpoint over the default one.
func WithEndpoint(endpoint string) ClientOption {
	return func(c *Client) {
		c.endpoint = endpoint
	}
}

// WithOAuthClientID will make the client use the given OAuth client ID.
func WithOAuthClientID(clientID string) ClientOption {
	return func(c *Client) {
		c.oAuthClientID = clientID
	}
}

// WithHTTPClient will make the client execute requests with the given [http.Client].
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.client = client
	}
}

// WithBleemeoAccountHeader will make the client include the given account ID
// in the X-Bleemeo-Account request header.
func WithBleemeoAccountHeader(accountID string) ClientOption {
	return func(c *Client) {
		c.customHeaders["X-Bleemeo-Account"] = accountID
	}
}
