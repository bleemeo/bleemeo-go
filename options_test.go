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
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"unsafe"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"golang.org/x/oauth2"
)

const oauthMockResponse = `HTTP/1.1 200 OK

{"access_token": "access", "expires_in": 36000, "token_type":` +
	` "Bearer", "scope": "read write", "refresh_token": "refresh"}`

type oauthMockTransport struct{}

func (omt oauthMockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return http.ReadResponse(bufio.NewReader(strings.NewReader(oauthMockResponse)), req) //nolint:wrapcheck
}

func mustParseURL(t *testing.T, s string) *url.URL {
	t.Helper()

	u, err := url.Parse(s)
	if err != nil {
		t.Fatalf("Could not parse URL %q: %v", s, err)
	}

	return u
}

func TestOptions(t *testing.T) {
	t.Parallel()

	creds := WithCredentials("u", "")
	oauthMockClient := &http.Client{Transport: oauthMockTransport{}}
	defaultEndpointURL := mustParseURL(t, defaultEndpoint)
	newOAuthTkCb := func(*oauth2.Token) {}

	cases := []struct {
		name           string
		env            map[string]string
		options        []ClientOption
		expectedError  error
		expectedClient Client
	}{
		{
			name:          "no options",
			expectedError: ErrNoAuthMeanProvided,
		},
		{
			name:    "no (optional) options",
			options: []ClientOption{creds},
			expectedClient: Client{
				username:      "u",
				endpoint:      defaultEndpoint,
				oAuthClientID: defaultOAuthClientID,
				client:        oauthMockClient,
				headers:       map[string]string{"User-Agent": defaultUserAgent},
				epURL:         defaultEndpointURL,
			},
		},
		{
			name:    "with credentials",
			options: []ClientOption{WithCredentials("usr", "pwd")},
			expectedClient: Client{
				username:      "usr",
				password:      "pwd",
				endpoint:      defaultEndpoint,
				oAuthClientID: defaultOAuthClientID,
				client:        oauthMockClient,
				headers:       map[string]string{"User-Agent": defaultUserAgent},
				epURL:         defaultEndpointURL,
			},
		},
		{
			name:    "with endpoint",
			options: []ClientOption{WithEndpoint("http://my-proxy.internal"), creds},
			expectedClient: Client{
				username:      "u",
				endpoint:      "http://my-proxy.internal",
				oAuthClientID: defaultOAuthClientID,
				client:        oauthMockClient,
				headers:       map[string]string{"User-Agent": defaultUserAgent},
				epURL:         mustParseURL(t, "http://my-proxy.internal"),
			},
		},
		{
			name:    "invalid endpoint",
			options: []ClientOption{WithEndpoint(":"), creds},
			expectedError: fmt.Errorf("invalid endpoint URL: %w", &url.Error{
				Op:  "parse",
				URL: ":",
				Err: errors.New("missing protocol scheme"), //nolint:err113
			}),
		},
		{
			name:    "with OAuth client ID",
			options: []ClientOption{WithOAuthClient("123456789", "53CR37"), creds},
			expectedClient: Client{
				username:          "u",
				endpoint:          defaultEndpoint,
				oAuthClientID:     "123456789",
				oAuthClientSecret: "53CR37",
				client:            oauthMockClient,
				headers:           map[string]string{"User-Agent": defaultUserAgent},
				epURL:             defaultEndpointURL,
			},
		},
		{
			name:    "with Bleemeo account header",
			options: []ClientOption{WithBleemeoAccountHeader("eea5c1dd-2edf-47b2-9ef6-7b239e16a5c3"), creds},
			expectedClient: Client{
				username:      "u",
				endpoint:      defaultEndpoint,
				oAuthClientID: defaultOAuthClientID,
				client:        oauthMockClient,
				headers: map[string]string{
					"User-Agent":        defaultUserAgent,
					"X-Bleemeo-Account": "eea5c1dd-2edf-47b2-9ef6-7b239e16a5c3",
				},
				epURL: defaultEndpointURL,
			},
		},
		{
			name:    "with initial OAuth refresh token",
			options: []ClientOption{WithInitialOAuthRefreshToken("initial")},
			expectedClient: Client{
				endpoint:            defaultEndpoint,
				oAuthClientID:       defaultOAuthClientID,
				oAuthInitialRefresh: "initial",
				client:              oauthMockClient,
				headers:             map[string]string{"User-Agent": defaultUserAgent},
				epURL:               defaultEndpointURL,
			},
		},
		{
			name: "with configuration from environment",
			env: map[string]string{
				"BLEEMEO_USER":                        "u",
				"BLEEMEO_PASSWORD":                    "p",
				"BLEEMEO_API_URL":                     "http://my-proxy.internal",
				"BLEEMEO_OAUTH_CLIENT_ID":             "123456789",
				"BLEEMEO_OAUTH_CLIENT_SECRET":         "53CR37",
				"BLEEMEO_ACCOUNT_ID":                  "eea5c1dd-2edf-47b2-9ef6-7b239e16a5c3",
				"BLEEMEO_OAUTH_INITIAL_REFRESH_TOKEN": "refresh-token",
			},
			options: []ClientOption{WithConfigurationFromEnv()},
			expectedClient: Client{
				username:            "u",
				password:            "p",
				endpoint:            "http://my-proxy.internal",
				oAuthClientID:       "123456789",
				oAuthClientSecret:   "53CR37",
				oAuthInitialRefresh: "refresh-token",
				client:              oauthMockClient,
				headers: map[string]string{
					"User-Agent":        defaultUserAgent,
					"X-Bleemeo-Account": "eea5c1dd-2edf-47b2-9ef6-7b239e16a5c3",
				},
				epURL: mustParseURL(t, "http://my-proxy.internal"),
			},
		},
		{
			name:    "with new oAuth token callback",
			options: []ClientOption{WithNewOAuthTokenCallback(newOAuthTkCb), creds},
			expectedClient: Client{
				username:              "u",
				endpoint:              defaultEndpoint,
				oAuthClientID:         defaultOAuthClientID,
				client:                oauthMockClient,
				newOAuthTokenCallback: newOAuthTkCb,
				headers:               map[string]string{"User-Agent": defaultUserAgent},
				epURL:                 defaultEndpointURL,
			},
		},
		// We can assume that WithHTTPClient() works since it is used in all the above cases.
	}

	for _, testCase := range cases {
		tc := testCase

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			for k, v := range tc.env {
				err := os.Setenv(k, v)
				if err != nil {
					t.Fatal("Failed to define environment:", err)
				}
			}

			client, err := NewClient(append(tc.options, WithHTTPClient(oauthMockClient))...)
			if diff := cmp.Diff(tc.expectedError, err, cmp.Comparer(errorComparer)); diff != "" {
				t.Fatalf("Client initialization: unexpected error (-want +got):\n%s", diff)
			}

			if err != nil {
				return
			}

			cmpOpts := cmp.Options{
				cmp.AllowUnexported(Client{}),
				cmpopts.IgnoreFields(Client{}, "authProvider"),
				cmp.Comparer(tokenCallbackComparer),
			}
			if diff := cmp.Diff(tc.expectedClient, *client, cmpOpts); diff != "" {
				t.Fatalf("Unexpected client: (-want +got)\n%s", diff)
			}
		})
	}
}

func errorComparer(x, y error) bool {
	return x.Error() == y.Error()
}

func tokenCallbackComparer(x, y func(*oauth2.Token)) bool {
	px := *(*unsafe.Pointer)(unsafe.Pointer(&x))
	py := *(*unsafe.Pointer)(unsafe.Pointer(&y))

	return px == py
}
