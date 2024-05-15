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
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"golang.org/x/oauth2"
)

func makeClientMockForAuth(
	t *testing.T, authHandler, revokeHandler mockHandler, expectedAccessTk *string, extraOpts ...ClientOption,
) (c *Client, requestCounter map[string]int, err error) {
	t.Helper()

	requestCounter = make(map[string]int)
	clientMock := &http.Client{
		Transport: &transportMock{
			handlers: map[string]mockHandler{
				tokenPath:          authHandler,
				"/o/revoke_token/": revokeHandler,
				"/v1/agent/<id>/": func(r *http.Request) (statusCode int, body []byte, err error) {
					authHeader := r.Header.Get("Authorization")

					if !strings.HasPrefix(authHeader, "Bearer ") {
						t.Fatalf("Invalid Authorization header: %q", authHeader)
					}

					token := strings.TrimPrefix(authHeader, "Bearer ")
					if token != *expectedAccessTk {
						// t.Fatalf("Unexpected access token: want %q, got %q", *expectedAccessTk, token)
						return http.StatusUnauthorized, []byte(`{}`), nil
					}

					return http.StatusOK, []byte(`{}`), nil
				},
			},
			counters: requestCounter,
		},
	}

	client, err := NewClient(append([]ClientOption{WithHTTPClient(clientMock)}, extraOpts...)...)

	return client, requestCounter, err
}

func TestAuthentication(t *testing.T) {
	t.Parallel()

	t.Run("with credentials", func(t *testing.T) {
		t.Parallel()

		const (
			username, password, clientID, clientSecret = "user", "passwd", "clID", "clScrt"
			firstAccessTk, firstRefreshTk              = "a-1", "r-1"
			secondAccessTk, secondRefreshTk            = "a-2", "r-2"
			tokenValidity                              = time.Hour
		)

		revokeCallsCount := 0

		authHandler := func(r *http.Request) (statusCode int, body []byte, err error) {
			reqBody, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal("Failed to read request body:", err)
			}

			reqBodyValues, err := url.ParseQuery(string(reqBody))
			if err != nil {
				t.Fatal("Failed to parse request body:", err)
			}

			var (
				expectedReqBody           url.Values
				accessToken, refreshToken string
			)

			switch reqBodyValues.Get("grant_type") {
			case "password":
				expectedReqBody = url.Values{
					"client_id":     {clientID},
					"client_secret": {clientSecret},
					"grant_type":    {"password"},
					"password":      {password},
					"username":      {username},
				}

				tokens, ok := map[int][2]string{
					0: {firstAccessTk, firstRefreshTk},
					1: {secondAccessTk, secondRefreshTk},
				}[revokeCallsCount]
				if !ok {
					t.Fatalf("Unexpected revoke calls count: %d (expected 2 max)", revokeCallsCount+1)
				}

				accessToken, refreshToken = tokens[0], tokens[1]
			case "refresh_token":
				expectedReqBody = url.Values{
					"client_id":     {clientID},
					"client_secret": {clientSecret},
					"grant_type":    {"refresh_token"},
					"refresh_token": {firstRefreshTk},
				}
				accessToken = secondAccessTk
				refreshToken = secondRefreshTk
			default:
				t.Fatalf("Unexpected grant type %v", reqBodyValues["grant_type"])
			}

			if diff := cmp.Diff(expectedReqBody, reqBodyValues); diff != "" {
				t.Fatalf("Unexpected request body (-want +got):\n%s", diff)
			}

			return http.StatusOK, []byte(fmt.Sprintf(
				`{"access_token": "%s", "expires_in": %d, "token_type": "Bearer", "scope": "read write", "refresh_token": "%s"}`,
				accessToken, int(tokenValidity.Seconds()), refreshToken)), nil
		}
		revokeHandler := func(*http.Request) (statusCode int, body []byte, err error) {
			revokeCallsCount++

			return http.StatusOK, nil, nil
		}
		expectedAccessToken := firstAccessTk

		client, requestCounter, err := makeClientMockForAuth(
			t, authHandler, revokeHandler, &expectedAccessToken,
			WithCredentials(username, password), WithOAuthClient(clientID, clientSecret),
		)
		if err != nil {
			t.Fatal("Failed to init client:", err)
		}

		token, err := client.authProvider.Token(context.Background())
		if err != nil {
			t.Fatal("Failed to retrieve token:", err)
		}

		expectedToken := &oauth2.Token{
			AccessToken:  firstAccessTk,
			TokenType:    "Bearer",
			RefreshToken: firstRefreshTk,
			Expiry:       time.Now().Add(tokenValidity),
		}
		cmpOpts := cmp.Options{cmpopts.IgnoreUnexported(oauth2.Token{}), cmpopts.EquateApproxTime(time.Minute)}

		if diff := cmp.Diff(expectedToken, token, cmpOpts); diff != "" {
			t.Fatalf("Unexpected token (-want +got):\n%s", diff)
		}

		_, err = client.Get(context.Background(), ResourceAgent, "<id>")
		if err != nil {
			t.Fatal("Failed to execute request:", err)
		}

		// Simulating a loss of validity from the token
		err = client.Logout(context.Background())
		if err != nil {
			t.Fatal("Failed to logout:", err)
		}

		expectedAccessToken = secondAccessTk

		_, err = client.Get(context.Background(), ResourceAgent, "<id>")
		if err != nil {
			t.Fatal("Failed to execute request:", err)
		}

		err = client.Logout(context.Background())
		if err != nil {
			t.Fatal("Failed to logout:", err)
		}

		expectedRequests := map[string]int{
			tokenPath:          2,
			"/o/revoke_token/": 2,
			"/v1/agent/<id>/":  2,
		}
		if diff := cmp.Diff(expectedRequests, requestCounter); diff != "" {
			t.Fatalf("Unexpected requests:\n%s", diff)
		}
	})

	t.Run("with initial refresh token", func(t *testing.T) {
		t.Parallel()

		const (
			initialRefresh, clientID, clientSecret = "refresh-init", "clID", "clScrt"
			firstAccessTk, firstRefreshTk          = "a-1", "r-1"
			tokenValidity                          = 10 * time.Hour
		)

		authHandler := func(r *http.Request) (statusCode int, body []byte, err error) {
			reqBody, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal("Failed to read request body:", err)
			}

			reqBodyValues, err := url.ParseQuery(string(reqBody))
			if err != nil {
				t.Fatal("Failed to parse request body:", err)
			}

			var (
				expectedReqBody           url.Values
				accessToken, refreshToken string
			)

			switch reqBodyValues.Get("grant_type") {
			case "refresh_token":
				expectedReqBody = url.Values{
					"client_id":     {clientID},
					"client_secret": {clientSecret},
					"grant_type":    {"refresh_token"},
					"refresh_token": {initialRefresh},
				}
				accessToken = firstAccessTk
				refreshToken = firstRefreshTk
			default:
				t.Fatalf("Unexpected grant type %v", reqBodyValues["grant_type"])
			}

			if diff := cmp.Diff(expectedReqBody, reqBodyValues); diff != "" {
				t.Fatalf("Unexpected request body (-want +got):\n%s", diff)
			}

			return http.StatusOK, []byte(fmt.Sprintf(
				`{"access_token": "%s", "expires_in": %d, "token_type": "Bearer", "scope": "read write", "refresh_token": "%s"}`,
				accessToken, int(tokenValidity.Seconds()), refreshToken)), nil
		}
		revokeHandler := func(*http.Request) (statusCode int, body []byte, err error) {
			return http.StatusOK, nil, nil
		}
		expectedAccessToken := firstAccessTk

		client, requestCounter, err := makeClientMockForAuth(
			t, authHandler, revokeHandler, &expectedAccessToken,
			WithInitialOAuthRefreshToken(initialRefresh), WithOAuthClient(clientID, clientSecret),
		)
		if err != nil {
			t.Fatal("Failed to init client:", err)
		}

		token, err := client.authProvider.Token(context.Background())
		if err != nil {
			t.Fatal("Failed to retrieve token:", err)
		}

		expectedToken := &oauth2.Token{
			AccessToken:  firstAccessTk,
			TokenType:    "Bearer",
			RefreshToken: firstRefreshTk,
			Expiry:       time.Now().Add(tokenValidity),
		}
		cmpOpts := cmp.Options{cmpopts.IgnoreUnexported(oauth2.Token{}), cmpopts.EquateApproxTime(time.Minute)}

		if diff := cmp.Diff(expectedToken, token, cmpOpts); diff != "" {
			t.Fatalf("Unexpected token (-want +got):\n%s", diff)
		}

		_, err = client.Get(context.Background(), ResourceAgent, "<id>")
		if err != nil {
			t.Fatal("Failed to execute request:", err)
		}

		err = client.authProvider.refetchToken(context.Background())
		if !errors.Is(err, ErrTokenIsRefreshOnly) {
			t.Fatalf("Expected refetch to fail with %v, got %v:", ErrTokenIsRefreshOnly, err)
		}

		err = client.Logout(context.Background())
		if err != nil {
			t.Fatal("Failed to logout:", err)
		}

		expectedRequests := map[string]int{
			tokenPath:          1,
			"/o/revoke_token/": 1,
			"/v1/agent/<id>/":  1,
		}
		if diff := cmp.Diff(expectedRequests, requestCounter); diff != "" {
			t.Fatalf("Unexpected requests:\n%s", diff)
		}
	})

	t.Run("with nothing", func(t *testing.T) {
		t.Parallel()

		_, err := NewClient()
		if !errors.Is(err, ErrNoAuthMeanProvided) {
			t.Fatalf("Expected error to be %v, got %v", ErrNoAuthMeanProvided, err)
		}
	})

	t.Run("invalid credentials", func(t *testing.T) {
		t.Parallel()

		respData := []byte(`{"error": "invalid_grant", "error_description": "Invalid credentials given."}`)
		authHandler := func(*http.Request) (statusCode int, body []byte, err error) {
			return http.StatusBadRequest, respData, nil
		}
		revokeHandler := func(*http.Request) (statusCode int, body []byte, err error) {
			return http.StatusOK, nil, nil
		}

		client, requestCounter, err := makeClientMockForAuth(
			t, authHandler, revokeHandler, new(string), WithCredentials("bad", "creds"),
		)
		if err != nil {
			t.Fatal("Unexpected error:", err)
		}

		_, err = client.Get(context.Background(), ResourceAgent, "<id>")
		if err == nil {
			t.Fatal("Expected an error, got none.")
		}

		unwrappErr, ok := err.(interface{ Unwrap() error })
		if !ok {
			t.Fatalf("Can't unwrap error of type %T", err)
		}

		expectedErr := &AuthError{
			APIError: &APIError{
				ReqPath:     "/v1/agent/<id>/",
				StatusCode:  400,
				ContentType: "",
				Message:     "Invalid credentials given.",
				Err: &oauth2.RetrieveError{
					ErrorCode:        "invalid_grant",
					ErrorDescription: "Invalid credentials given.",
					Body:             respData,
				},
				Response: respData,
			},
			ErrorCode: "invalid_grant",
		}
		cmpOpts := cmp.Options{cmpopts.IgnoreFields(oauth2.RetrieveError{}, "Response")}

		if diff := cmp.Diff(expectedErr, unwrappErr.Unwrap(), cmpOpts); diff != "" {
			t.Fatalf("Unexpected token (-want +got):\n%s", diff)
		}

		expectedRequests := map[string]int{
			tokenPath: 1,
		}
		if diff := cmp.Diff(expectedRequests, requestCounter); diff != "" {
			t.Fatalf("Unexpected requests:\n%s", diff)
		}
	})

	t.Run("auth header injection", func(t *testing.T) {
		t.Parallel()

		t.Run("valid token", func(t *testing.T) {
			t.Parallel()

			const tokenType, accessTk = "Bearer", "access"

			ap := &authenticationProvider{
				newToken: func(context.Context) (*oauth2.Token, error) {
					return &oauth2.Token{
						AccessToken: accessTk,
						TokenType:   tokenType,
					}, nil
				},
			}

			req, err := http.NewRequest(http.MethodGet, "http://localhost/route", nil) //nolint: noctx
			if err != nil {
				t.Fatal("Can't make request:", err)
			}

			err = ap.injectHeader(context.Background(), req)
			if err != nil {
				t.Fatal("Failed to inject auth header:", err)
			}

			authHeader := req.Header.Get("Authorization")
			expectedAuthHeader := tokenType + " " + accessTk

			if authHeader != expectedAuthHeader {
				t.Fatalf("Unexpected auth header: want %q, got %q", expectedAuthHeader, authHeader)
			}
		})

		t.Run("invalid token", func(t *testing.T) {
			t.Parallel()

			tokenRetErr := &oauth2.RetrieveError{
				ErrorCode:        "error code",
				ErrorDescription: "error desc",
				Response: &http.Response{
					StatusCode: http.StatusBadRequest,
					Header:     map[string][]string{"Content-Type": {"application/json"}},
				},
			}
			ap := &authenticationProvider{
				newToken: func(context.Context) (*oauth2.Token, error) {
					return nil, tokenRetErr
				},
			}

			req, err := http.NewRequest(http.MethodGet, "http://localhost/route", nil) //nolint: noctx
			if err != nil {
				t.Fatal("Can't make request:", err)
			}

			err = ap.injectHeader(context.Background(), req)
			if err == nil {
				t.Fatal("Expected error, got none")
			}

			expectedErr := &AuthError{
				APIError: &APIError{
					ReqPath:     "/route",
					StatusCode:  400,
					ContentType: "application/json",
					Message:     "error desc",
					Err:         tokenRetErr,
				},
				ErrorCode: "error code",
			}
			if diff := cmp.Diff(expectedErr, err); diff != "" {
				t.Fatalf("Unexpected error (-want +got):\n%s", diff)
			}
		})
	})
}
