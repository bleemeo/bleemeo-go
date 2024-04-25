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

type ClientOption func(*Client)

func WithCredentials(username, password string) ClientOption {
	return func(c *Client) {
		c.username = username
		c.password = password
	}
}

func WithCredentialsFromEnv() ClientOption {
	return func(c *Client) {
		c.username = os.Getenv("BLEEMEO_USER")
		c.password = os.Getenv("BLEEMEO_PASSWORD")

	}
}

func WithEndpoint(endpoint string) ClientOption {
	return func(c *Client) {
		c.endpoint = endpoint
	}
}

func WithOAuthClientID(clientID string) ClientOption {
	return func(c *Client) {
		c.oAuthClientID = clientID
	}
}

func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.client = client
	}
}

func WithBleemeoAccountHeader(accountID string) ClientOption {
	return func(c *Client) {
		c.customHeaders["X-Bleemeo-Account"] = accountID
	}
}
