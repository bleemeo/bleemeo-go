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
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type authenticationProvider struct {
	tokenSource oauth2.TokenSource
}

// newCredentialsAuthProvider makes a new token source based on the given credentials.
// New tokens will be fetched with the "password" grant type.
func newCredentialsAuthProvider(endpointURL, username, password, clientID, clientSecret string, client *http.Client) authenticationProvider {
	cfg := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     endpointURL + "/o/token/",
		EndpointParams: url.Values{
			"grant_type": {"password"},
			"username":   {username},
			"password":   {password},
		},
		AuthStyle: oauth2.AuthStyleInParams,
	}

	return authenticationProvider{
		tokenSource: cfg.TokenSource(context.WithValue(context.Background(), oauth2.HTTPClient, client)),
	}
}

// newRefreshAuthProvider makes a new token source based on the given refresh token.
// New tokens will be fetched with the "refresh_token" grant type.
func newRefreshAuthProvider(endpointURL, clientID, refreshToken string, client *http.Client) authenticationProvider {
	cfg := oauth2.Config{
		ClientID: clientID,
		Endpoint: oauth2.Endpoint{
			TokenURL:  endpointURL + "/o/token/",
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}
	initialToken := oauth2.Token{
		RefreshToken: refreshToken,
	}

	return authenticationProvider{
		tokenSource: cfg.TokenSource(context.WithValue(context.Background(), oauth2.HTTPClient, client), &initialToken),
	}
}

func (ap authenticationProvider) token() (*oauth2.Token, error) {
	return ap.tokenSource.Token()
}

func (ap authenticationProvider) injectHeader(req *http.Request) error {
	tk, err := ap.tokenSource.Token()
	if err != nil {
		if retErr := new(oauth2.RetrieveError); errors.As(err, &retErr) {
			return &AuthError{
				ClientError: ClientError{
					apiError: apiError{
						ReqPath:     req.URL.Path,
						StatusCode:  retErr.Response.StatusCode,
						ContentType: retErr.Response.Header.Get("Content-Type"),
						Message:     retErr.ErrorDescription,
						Err:         err,
						Response:    retErr.Body,
					},
				},
				ErrorCode: retErr.ErrorCode,
			}
		}

		return err
	}

	tk.SetAuthHeader(req)

	return nil
}
