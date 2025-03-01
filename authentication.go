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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const tokenPath = "/o/token/"

type (
	tokenProvider  func(ctx context.Context) (*oauth2.Token, error)
	tokenRefresher func(ctx context.Context, refreshToken string) (*oauth2.Token, error)
)

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

func newRefresher(endpointURL *url.URL, clientID, clientSecret string, client *http.Client) tokenRefresher {
	cfg := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL:  endpointURL.JoinPath(tokenPath).String(),
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
	refreshOnly            bool
	clientID, clientSecret string
	newOAuthTokenCallback  func(token *oauth2.Token)

	httpClient   *http.Client
	newToken     tokenProvider
	refreshToken tokenRefresher
	token        *oauth2.Token
}

func newAuthenticationProvider(
	endpointURL *url.URL, username, password, initialRefreshToken, clientID, clientSecret string, client *http.Client,
) *authenticationProvider {
	client = wrapTransportWithUserAgent(client, defaultUserAgent)
	authProvider := authenticationProvider{
		clientID:     clientID,
		clientSecret: clientSecret,
		httpClient:   client,
		refreshToken: newRefresher(endpointURL, clientID, clientSecret, client),
	}

	if username != "" {
		authProvider.refreshOnly = false
		authProvider.newToken = credentialsTokenProvider(endpointURL, username, password, clientID, clientSecret, client)
	} else {
		authProvider.refreshOnly = true
	}

	if initialRefreshToken != "" {
		authProvider.token = &oauth2.Token{
			RefreshToken: initialRefreshToken,
		}
	}

	return &authProvider
}

// newCredentialsAuthProvider makes a new token source based on the given credentials.
// New tokens will be fetched with the "password" grant type.
func credentialsTokenProvider(
	endpointURL *url.URL, username, password, clientID, clientSecret string, client *http.Client,
) tokenProvider {
	cfg := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     endpointURL.JoinPath(tokenPath).String(),
		EndpointParams: url.Values{
			"grant_type": {"password"},
			"username":   {username},
			"password":   {password},
		},
		AuthStyle: oauth2.AuthStyleInParams,
	}
	client = wrapTransportWithUserAgent(client, defaultUserAgent)

	return func(ctx context.Context) (*oauth2.Token, error) {
		return cfg.TokenSource(context.WithValue(ctx, oauth2.HTTPClient, client)).Token()
	}
}

func (ap *authenticationProvider) Token(ctx context.Context) (*oauth2.Token, error) {
	ap.l.Lock()
	defer ap.l.Unlock()

	var err error

	switch {
	case ap.token == nil:
		if ap.newToken == nil {
			return nil, ErrNoAuthMeanProvided
		}

		ap.token, err = ap.newToken(ctx)
	case !ap.token.Valid():
		if ap.token.RefreshToken == "" {
			return nil, errTokenHasNoRefresh
		}

		ap.token, err = ap.refreshToken(ctx, ap.token.RefreshToken)
		if err != nil {
			if !ap.refreshOnly {
				ap.token, err = ap.newToken(ctx)
			}
		}
	default:
		return ap.token, nil
	}

	// A new token has been retrieved (if no error occurred)
	if ap.newOAuthTokenCallback != nil && err == nil {
		ap.newOAuthTokenCallback(ap.token)
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

	if ap.newOAuthTokenCallback != nil {
		ap.newOAuthTokenCallback(ap.token)
	}

	return nil
}

func (ap *authenticationProvider) injectHeader(ctx context.Context, req *http.Request) error {
	tk, err := ap.Token(ctx)
	if err != nil {
		if retErr := new(oauth2.RetrieveError); errors.As(err, &retErr) {
			return buildAuthError(req.URL.Path, retErr)
		}

		return fmt.Errorf("failed to retrieve authentication token: %w", err)
	}

	tk.SetAuthHeader(req)

	return nil
}

func (ap *authenticationProvider) logout(ctx context.Context, endpoint string) error {
	ap.l.Lock()
	defer ap.l.Unlock()

	if ap.token == nil || !ap.token.Valid() {
		return nil // No need to perform a logout
	}

	endpointURL, err := url.Parse(endpoint)
	if err != nil {
		return fmt.Errorf("can't parse endpoint URL: %w", err)
	}

	reqURL, err := endpointURL.Parse("/o/revoke_token/")
	if err != nil {
		return fmt.Errorf("can't parse logout URL: %w", err)
	}

	values := url.Values{
		"client_id":       {ap.clientID},
		"token_type_hint": {"refresh_token"},
		"token":           {ap.token.RefreshToken},
	}

	if ap.clientSecret != "" {
		values.Set("client_secret", ap.clientSecret)
	}

	// Revoking the refresh token will also revoke the related access token
	body := strings.NewReader(values.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL.String(), body)
	if err != nil {
		return fmt.Errorf("failed to parse logout request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := ap.httpClient.Do(req)
	if err != nil {
		// Multiple error verbs are only possible since Go1.20
		return fmt.Errorf("%s: %w", ErrTokenRevoke.Error(), err)
	}

	defer cleanupResponse(resp)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: server replyed with status code %d", ErrTokenRevoke, resp.StatusCode)
	}

	ap.token = nil

	return nil
}

func buildAuthError(reqPath string, retErr *oauth2.RetrieveError) *AuthError {
	return &AuthError{
		APIError: &APIError{
			ReqPath:     reqPath,
			StatusCode:  retErr.Response.StatusCode,
			ContentType: retErr.Response.Header.Get("Content-Type"),
			Message:     retErr.ErrorDescription,
			Err:         retErr,
			Response:    retErr.Body,
		},
		ErrorCode: retErr.ErrorCode,
	}
}

func buildAuthErrorFromBody(apiErr *APIError) error {
	authErr := AuthError{
		APIError: apiErr,
	}

	var respData struct {
		// OAuth response
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
		// API response
		Detail   string `json:"detail"`
		Code     string `json:"code"`
		Messages []struct {
			Message string `json:"message"`
		} `json:"messages"`
	}

	if err := json.Unmarshal(apiErr.Response, &respData); err != nil {
		authErr.Err = &JSONUnmarshalError{
			&jsonError{
				Err:      err,
				DataKind: JsonErrorDataKind_401Details,
				Data:     apiErr.Response,
			},
		}

		return &authErr
	}

	switch {
	case respData.Error != "":
		authErr.ErrorCode = respData.Error
		authErr.Message = respData.ErrorDescription
	case len(respData.Messages) > 0:
		authErr.Message = respData.Messages[0].Message
	case respData.Detail != "":
		authErr.Message = respData.Detail
	default:
		authErr.Message = string(apiErr.Response)
	}

	return &authErr
}
