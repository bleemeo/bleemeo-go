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
	"time"

	"golang.org/x/oauth2"
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

// WithBleemeoAccountHeader will make the client include the given account ID
// in the X-Bleemeo-Account request header.
// This is required if your credentials have access to multiple Bleemeo accounts.
func WithBleemeoAccountHeader(accountID string) ClientOption {
	return func(c *Client) {
		c.headers["X-Bleemeo-Account"] = accountID
	}
}

// WithOAuthClient will make the client use the given OAuth client ID/secret over the default one.
func WithOAuthClient(clientID, clientSecret string) ClientOption {
	return func(c *Client) {
		c.oAuthClientID = clientID
		c.oAuthClientSecret = clientSecret
	}
}

// WithEndpoint will make the client use the given endpoint over the default one.
func WithEndpoint(endpoint string) ClientOption {
	return func(c *Client) {
		c.endpoint = endpoint
	}
}

// WithInitialOAuthRefreshToken will make the client prefer the given refresh token
// over regular user/password credentials for authenticating against the API.
func WithInitialOAuthRefreshToken(refreshToken string) ClientOption {
	return func(c *Client) {
		c.oAuthInitialRefresh = refreshToken
	}
}

// WithConfigurationFromEnv will make the client retrieve and use configuration options
// defined in environment variables, such as
//
// - credentials: "BLEEMEO_USER" & "BLEEMEO_PASSWORD"
//
// - Bleemeo account ID: "BLEEMEO_ACCOUNT_ID"
//
// - OAuth client ID/secret: "BLEEMEO_OAUTH_CLIENT_ID" & "BLEEMEO_OAUTH_CLIENT_SECRET"
//
// - API URL: "BLEEMEO_API_URL"
//
// - Initial refresh token: "BLEEMEO_OAUTH_INITIAL_REFRESH_TOKEN".
func WithConfigurationFromEnv() ClientOption {
	return func(c *Client) {
		if username, set := os.LookupEnv("BLEEMEO_USER"); set {
			c.username = username
		}

		if password, set := os.LookupEnv("BLEEMEO_PASSWORD"); set {
			c.password = password
		}

		if accountID, set := os.LookupEnv("BLEEMEO_ACCOUNT_ID"); set {
			c.headers["X-Bleemeo-Account"] = accountID
		}

		if oAuthClientID, set := os.LookupEnv("BLEEMEO_OAUTH_CLIENT_ID"); set {
			c.oAuthClientID = oAuthClientID
		}

		if apiURL, set := os.LookupEnv("BLEEMEO_API_URL"); set {
			c.endpoint = apiURL
		}

		if oAuthClientSecret, set := os.LookupEnv("BLEEMEO_OAUTH_CLIENT_SECRET"); set {
			c.oAuthClientSecret = oAuthClientSecret
		}

		if refreshToken, set := os.LookupEnv("BLEEMEO_OAUTH_INITIAL_REFRESH_TOKEN"); set {
			c.oAuthInitialRefresh = refreshToken
		}
	}
}

// WithHTTPClient will make the client execute requests with the given [http.Client].
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.client = client
	}
}

// WithNewOAuthTokenCallback defines the given function as the one to be called
// when a new OAuth token is retrieved.
func WithNewOAuthTokenCallback(callback func(token *oauth2.Token)) ClientOption {
	return func(c *Client) {
		c.newOAuthTokenCallback = callback
	}
}

// WithThrottleMaxDelayAutoRetry defines the delay under which throttled requests are automatically resent.
func WithThrottleMaxDelayAutoRetry(delay time.Duration) ClientOption {
	return func(c *Client) {
		c.throttleMaxAutoRetryDelay = delay
	}
}
