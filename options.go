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

// WithEndpoint will make the client use the given endpoint over the default one.
func WithEndpoint(endpoint string) ClientOption {
	return func(c *Client) {
		c.endpoint = endpoint
	}
}

// WithOAuthClientID will make the client use the given OAuth client ID/secret.
func WithOAuthClientID(clientID, clientSecret string) ClientOption {
	return func(c *Client) {
		c.oAuthClientID = clientID
		c.oAuthClientSecret = clientSecret
	}
}

// WithBleemeoAccountHeader will make the client include the given account ID
// in the X-Bleemeo-Account request header.
func WithBleemeoAccountHeader(accountID string) ClientOption {
	return func(c *Client) {
		c.customHeaders["X-Bleemeo-Account"] = accountID
	}
}

// WithConfigurationFromEnv will make the client retrieve and use configuration options
// defined in environment variables, such as
//
// - credentials: "BLEEMEO_USER" & "BLEEMEO_PASSWORD"
//
// - API URL: "BLEEMEO_API_URL"
//
// - OAuth client ID/secret: "BLEEMEO_OAUTH_CLIENT_ID" & "BLEEMEO_OAUTH_CLIENT_SECRET"
//
// - Bleemeo account ID: "BLEEMEO_ACCOUNT_ID"
func WithConfigurationFromEnv() ClientOption {
	return func(c *Client) {
		if username, set := os.LookupEnv("BLEEMEO_USER"); set {
			c.username = username
		}

		if password, set := os.LookupEnv("BLEEMEO_PASSWORD"); set {
			c.password = password
		}

		if apiURL, set := os.LookupEnv("BLEEMEO_API_URL"); set {
			c.endpoint = apiURL
		}

		if oAuthClientID, set := os.LookupEnv("BLEEMEO_OAUTH_CLIENT_ID"); set {
			c.oAuthClientID = oAuthClientID
		}

		if oAuthClientSecret, set := os.LookupEnv("BLEEMEO_OAUTH_CLIENT_SECRET"); set {
			c.oAuthClientSecret = oAuthClientSecret
		}

		if accountID, set := os.LookupEnv("BLEEMEO_ACCOUNT_ID"); set {
			c.customHeaders["X-Bleemeo-Account"] = accountID
		}
	}
}

// WithHTTPClient will make the client execute requests with the given [http.Client].
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.client = client
	}
}
