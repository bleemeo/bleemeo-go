package bleemeo

import (
	"context"
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

func makeClientMockForAuth(t *testing.T, authHandler mockHandler, expectedAccessTk string, extraOpts ...ClientOption) (c *Client, requestCounter map[string]int, err error) {
	t.Helper()

	requestCounter = make(map[string]int)
	clientMock := &http.Client{
		Transport: &transportMock{
			handlers: map[string]mockHandler{
				"/o/token/": authHandler,
				"/o/revoke_token/": func(*http.Request) (statusCode int, body []byte, err error) {
					return http.StatusOK, nil, nil
				},
				"/v1/agent/<id>/": func(r *http.Request) (statusCode int, body []byte, err error) {
					authHeader := r.Header.Get("Authorization")

					if !strings.HasPrefix(authHeader, "Bearer ") {
						t.Fatalf("Invalid Authorization header: %q", authHeader)
					}

					token := strings.TrimPrefix(authHeader, "Bearer ")
					if token != expectedAccessTk {
						t.Fatalf("Invalid access token: %q", token)
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

type tokenSourceMock struct {
	token *oauth2.Token
	err   error
}

func (tsm tokenSourceMock) Token() (*oauth2.Token, error) {
	return tsm.token, tsm.err
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

		handler := func(r *http.Request) (statusCode int, body []byte, err error) {
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
				accessToken = firstAccessTk
				refreshToken = firstRefreshTk
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

			if diff := cmp.Diff(reqBodyValues, expectedReqBody); diff != "" {
				t.Fatalf("Unexpected request body (-want +got):\n%s", diff)
			}

			return http.StatusOK, []byte(fmt.Sprintf(`{"access_token": "%s", "expires_in": %d, "token_type": "Bearer", "scope": "read write", "refresh_token": "%s"}`, accessToken, int(tokenValidity.Seconds()), refreshToken)), nil
		}

		client, requestCounter, err := makeClientMockForAuth(t, handler, secondAccessTk, WithCredentials(username, password), WithOAuthClient(clientID, clientSecret))
		if err != nil {
			t.Fatal("Failed to init client:", err)
		}

		token, err := client.authProvider.Token()
		if err != nil {
			t.Fatal("Failed to retrieve token:", err)
		}

		expectedToken := &oauth2.Token{
			AccessToken:  secondAccessTk,
			TokenType:    "Bearer",
			RefreshToken: secondRefreshTk,
			Expiry:       time.Now().Add(tokenValidity),
		}
		if diff := cmp.Diff(token, expectedToken, cmpopts.IgnoreUnexported(oauth2.Token{}), cmpopts.EquateApproxTime(time.Minute)); diff != "" {
			t.Fatalf("Unexpected token (-want +got):\n%s", diff)
		}

		_, err = client.Get(context.Background(), ResourceAgent, "<id>", DefaultFields)
		if err != nil {
			t.Fatal("Failed to execute request:", err)
		}

		err = client.Logout(context.Background())
		if err != nil {
			t.Fatal("Failed to logout:", err)
		}

		expectedRequests := map[string]int{
			"/o/token/":        2,
			"/o/revoke_token/": 1,
			"/v1/agent/<id>/":  1,
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

		handler := func(r *http.Request) (statusCode int, body []byte, err error) {
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

			if diff := cmp.Diff(reqBodyValues, expectedReqBody); diff != "" {
				t.Fatalf("Unexpected request body (-want +got):\n%s", diff)
			}

			return http.StatusOK, []byte(fmt.Sprintf(`{"access_token": "%s", "expires_in": %d, "token_type": "Bearer", "scope": "read write", "refresh_token": "%s"}`, accessToken, int(tokenValidity.Seconds()), refreshToken)), nil
		}

		client, requestCounter, err := makeClientMockForAuth(t, handler, firstAccessTk, WithInitialOAuthRefreshToken(initialRefresh), WithOAuthClient(clientID, clientSecret))
		if err != nil {
			t.Fatal("Failed to init client:", err)
		}

		token, err := client.authProvider.Token()
		if err != nil {
			t.Fatal("Failed to retrieve token:", err)
		}

		expectedToken := &oauth2.Token{
			AccessToken:  firstAccessTk,
			TokenType:    "Bearer",
			RefreshToken: firstRefreshTk,
			Expiry:       time.Now().Add(tokenValidity),
		}
		if diff := cmp.Diff(token, expectedToken, cmpopts.IgnoreUnexported(oauth2.Token{}), cmpopts.EquateApproxTime(time.Minute)); diff != "" {
			t.Fatalf("Unexpected token (-want +got):\n%s", diff)
		}

		_, err = client.Get(context.Background(), ResourceAgent, "<id>", DefaultFields)
		if err != nil {
			t.Fatal("Failed to execute request:", err)
		}

		err = client.Logout(context.Background())
		if err != nil {
			t.Fatal("Failed to logout:", err)
		}

		expectedRequests := map[string]int{
			"/o/token/":        1,
			"/o/revoke_token/": 1,
			"/v1/agent/<id>/":  1,
		}
		if diff := cmp.Diff(expectedRequests, requestCounter); diff != "" {
			t.Fatalf("Unexpected requests:\n%s", diff)
		}
	})

	t.Run("invalid credentials", func(t *testing.T) {
		t.Parallel()

		respData := []byte(`{"error": "invalid_grant", "error_description": "Invalid credentials given."}`)
		handler := func(r *http.Request) (statusCode int, body []byte, err error) {
			return http.StatusBadRequest, respData, nil
		}

		_, requestCounter, err := makeClientMockForAuth(t, handler, "", WithCredentials("bad", "creds"), WithOAuthClient("id", ""))

		expectedErr := &AuthError{
			ClientError: ClientError{
				apiError{
					ReqPath:     "/o/token/",
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
			},
			ErrorCode: "invalid_grant",
		}
		if diff := cmp.Diff(err, expectedErr, cmp.AllowUnexported(ClientError{}), cmpopts.IgnoreFields(oauth2.RetrieveError{}, "Response")); diff != "" {
			t.Fatalf("Unexpected token (-want +got):\n%s", diff)
		}

		expectedRequests := map[string]int{
			"/o/token/": 1,
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
				tokenSource: tokenSourceMock{
					token: &oauth2.Token{
						AccessToken: accessTk,
						TokenType:   tokenType,
					},
					err: nil,
				},
			}

			req, err := http.NewRequest(http.MethodGet, "http://localhost/route", nil)
			if err != nil {
				t.Fatal("Can't make request:", err)
			}

			err = ap.injectHeader(req)
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
					StatusCode: 400,
					Header:     map[string][]string{"Content-Type": {"application/json"}},
				},
			}
			ap := &authenticationProvider{
				tokenSource: tokenSourceMock{
					token: nil,
					err:   tokenRetErr,
				},
			}

			req, err := http.NewRequest(http.MethodGet, "http://localhost/route", nil)
			if err != nil {
				t.Fatal("Can't make request:", err)
			}

			err = ap.injectHeader(req)
			if err == nil {
				t.Fatal("Expected error, got none")
			}

			expectedErr := &AuthError{
				ClientError: ClientError{
					apiError{
						ReqPath:     "/route",
						StatusCode:  400,
						ContentType: "application/json",
						Message:     "error desc",
						Err:         tokenRetErr,
					},
				},
				ErrorCode: "error code",
			}
			if diff := cmp.Diff(err, expectedErr, cmp.AllowUnexported(ClientError{})); diff != "" {
				t.Fatalf("Unexpected error (-want +got):\n%s", diff)
			}
		})
	})
}
