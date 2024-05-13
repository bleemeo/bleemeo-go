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
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const tokenPath = "/o/token/"

type userAgentTransporter struct {
	userAgentHeader string
	http.RoundTripper
}

func (t userAgentTransporter) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", t.userAgentHeader)

	return t.RoundTripper.RoundTrip(req) //nolint:wrapcheck
}

func wrapTransportWithUserAgent(client *http.Client, userAgentHeader string) *http.Client {
	c := *client // Avoid mutating the given client

	initialTransport := client.Transport
	if initialTransport == nil {
		initialTransport = http.DefaultTransport
	}

	c.Transport = userAgentTransporter{
		userAgentHeader: userAgentHeader,
		RoundTripper:    initialTransport,
	}

	return &c
}

type tokenRefresher func(ctx context.Context, refreshToken string) (*oauth2.Token, error)

func newRefresher(endpointURL, clientID, clientSecret string, client *http.Client) tokenRefresher {
	cfg := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL:  endpointURL + tokenPath,
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}

	return func(ctx context.Context, refreshToken string) (*oauth2.Token, error) {
		ctx = context.WithValue(ctx, oauth2.HTTPClient, client)
		refTk := oauth2.Token{
			RefreshToken: refreshToken,
		}

		return cfg.TokenSource(ctx, &refTk).Token()
	}
}

type authenticationProvider struct {
	l sync.Mutex
	// Whether this provider only supports token refresh or not.
	refreshOnly bool

	httpClient   *http.Client
	newToken     func(ctx context.Context) (*oauth2.Token, error)
	refreshToken tokenRefresher
	token        *oauth2.Token
}

// newCredentialsAuthProvider makes a new token source based on the given credentials.
// New tokens will be fetched with the "password" grant type.
func newCredentialsAuthProvider(endpointURL, username, password, clientID, clientSecret string, client *http.Client) *authenticationProvider {
	cfg := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     endpointURL + tokenPath,
		EndpointParams: url.Values{
			"grant_type": {"password"},
			"username":   {username},
			"password":   {password},
		},
		AuthStyle: oauth2.AuthStyleInParams,
	}
	client = wrapTransportWithUserAgent(client, defaultUserAgent)

	return &authenticationProvider{
		refreshOnly: false,
		httpClient:  client,
		newToken: func(ctx context.Context) (*oauth2.Token, error) {
			return cfg.TokenSource(context.WithValue(ctx, oauth2.HTTPClient, client)).Token()
		},
		refreshToken: newRefresher(endpointURL, clientID, clientSecret, client),
	}
}

// newRefreshAuthProvider makes a new token source based on the given refresh token.
// New tokens will be fetched with the "refresh_token" grant type.
func newRefreshAuthProvider(endpointURL, clientID, clientSecret, refreshToken string, client *http.Client) *authenticationProvider {
	client = wrapTransportWithUserAgent(client, defaultUserAgent)
	refresher := newRefresher(endpointURL, clientID, clientSecret, client)

	return &authenticationProvider{
		refreshOnly: true,
		httpClient:  client,
		newToken: func(ctx context.Context) (*oauth2.Token, error) {
			return refresher(ctx, refreshToken) // The first token is actually granted by a refresh
		},
		refreshToken: refresher,
	}
}

func (ap *authenticationProvider) Token() (*oauth2.Token, error) {
	ap.l.Lock()
	defer ap.l.Unlock()

	var err error

	switch {
	case ap.token == nil:
		ap.token, err = ap.newToken(context.Background())
	case !ap.token.Valid():
		if ap.token.RefreshToken == "" {
			return nil, ErrTokenHasNoRefresh
		}

		ap.token, err = ap.refreshToken(context.Background(), ap.token.RefreshToken)
		if err != nil {
			if !ap.refreshOnly {
				ap.token, err = ap.newToken(context.Background())
			}
		}
	}

	return ap.token, err
}

func (ap *authenticationProvider) refetchToken(ctx context.Context) error {
	if ap.refreshOnly {
		return ErrTokenIsRefreshOnly
	}

	ap.l.Lock()
	defer ap.l.Unlock()

	tk, err := ap.newToken(ctx)
	if err != nil {
		if retErr := new(oauth2.RetrieveError); errors.As(err, &retErr) {
			return buildAuthError(tokenPath, retErr)
		}

		return err
	}

	ap.token = tk

	return nil
}

func (ap *authenticationProvider) injectHeader(req *http.Request) error {
	tk, err := ap.Token()
	if err != nil {
		if retErr := new(oauth2.RetrieveError); errors.As(err, &retErr) {
			return buildAuthError(req.URL.Path, retErr)
		}

		return fmt.Errorf("failed to retrieve authentication token: %w", err)
	}

	tk.SetAuthHeader(req)

	return nil
}

func buildAuthError(reqPath string, retErr *oauth2.RetrieveError) *AuthError {
	return &AuthError{
		ClientError: ClientError{
			apiError: apiError{
				ReqPath:     reqPath,
				StatusCode:  retErr.Response.StatusCode,
				ContentType: retErr.Response.Header.Get("Content-Type"),
				Message:     retErr.ErrorDescription,
				Err:         retErr,
				Response:    retErr.Body,
			},
		},
		ErrorCode: retErr.ErrorCode,
	}
}
