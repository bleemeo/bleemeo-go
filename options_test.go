package bleemeo

import (
	"bufio"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

const oauthMockResponse = `HTTP/1.1 200 OK

{"access_token": "access", "expires_in": 36000, "token_type": "Bearer", "scope": "read write", "refresh_token": "refresh"}`

type oauthMockTransport struct{}

func (omt oauthMockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	rdr := bufio.NewReader(strings.NewReader(oauthMockResponse))
	return http.ReadResponse(rdr, req)
}

func mustParseURL(t *testing.T, s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		t.Fatalf("Could not parse URL %q: %v", s, err)
	}

	return u
}

func TestOptions(t *testing.T) {
	oauthMockClient := &http.Client{Transport: oauthMockTransport{}}
	defaultEndpointURL := mustParseURL(t, defaultEndpoint)
	oauthClientOpt := WithOAuthClient("id", "") // the client ID is mandatory

	cases := []struct {
		name           string
		env            map[string]string
		options        []ClientOption
		expectedError  error
		expectedClient Client
	}{
		{
			name:          "no options",
			expectedError: ErrNoOAuthClientIDProvided,
		},
		{
			name:    "with credentials",
			options: []ClientOption{WithCredentials("u", "p"), oauthClientOpt},
			expectedClient: Client{
				username:            "u",
				password:            "p",
				endpoint:            defaultEndpoint,
				oAuthClientID:       "id",
				oAuthInitialRefresh: "refresh",
				client:              oauthMockClient,
				customHeaders:       map[string]string{"User-Agent": defaultUserAgent},
				epURL:               defaultEndpointURL,
			},
		},
		{
			name:    "with endpoint",
			options: []ClientOption{WithEndpoint("http://my-proxy.internal"), oauthClientOpt},
			expectedClient: Client{
				endpoint:            "http://my-proxy.internal",
				oAuthClientID:       "id",
				oAuthInitialRefresh: "refresh",
				client:              oauthMockClient,
				customHeaders:       map[string]string{"User-Agent": defaultUserAgent},
				epURL:               mustParseURL(t, "http://my-proxy.internal"),
			},
		},
		{
			name:    "with OAuth client ID",
			options: []ClientOption{WithOAuthClient("123456789", "53CR37")},
			expectedClient: Client{
				endpoint:            defaultEndpoint,
				oAuthClientID:       "123456789",
				oAuthClientSecret:   "53CR37",
				oAuthInitialRefresh: "refresh",
				client:              oauthMockClient,
				customHeaders:       map[string]string{"User-Agent": defaultUserAgent},
				epURL:               defaultEndpointURL,
			},
		},
		{
			name:    "with Bleemeo account header",
			options: []ClientOption{WithBleemeoAccountHeader("eea5c1dd-2edf-47b2-9ef6-7b239e16a5c3"), oauthClientOpt},
			expectedClient: Client{
				endpoint:            defaultEndpoint,
				oAuthClientID:       "id",
				oAuthInitialRefresh: "refresh",
				client:              oauthMockClient,
				customHeaders:       map[string]string{"User-Agent": defaultUserAgent, "X-Bleemeo-Account": "eea5c1dd-2edf-47b2-9ef6-7b239e16a5c3"},
				epURL:               defaultEndpointURL,
			},
		},
		{
			name: "with configuration from environment",
			env: map[string]string{
				"BLEEMEO_USER":                "u",
				"BLEEMEO_PASSWORD":            "p",
				"BLEEMEO_API_URL":             "http://my-proxy.internal",
				"BLEEMEO_OAUTH_CLIENT_ID":     "123456789",
				"BLEEMEO_OAUTH_CLIENT_SECRET": "53CR37",
				"BLEEMEO_ACCOUNT_ID":          "eea5c1dd-2edf-47b2-9ef6-7b239e16a5c3",
			},
			options: []ClientOption{WithConfigurationFromEnv()},
			expectedClient: Client{
				username:            "u",
				password:            "p",
				endpoint:            "http://my-proxy.internal",
				oAuthClientID:       "123456789",
				oAuthClientSecret:   "53CR37",
				oAuthInitialRefresh: "refresh",
				client:              oauthMockClient,
				customHeaders:       map[string]string{"User-Agent": defaultUserAgent, "X-Bleemeo-Account": "eea5c1dd-2edf-47b2-9ef6-7b239e16a5c3"},
				epURL:               mustParseURL(t, "http://my-proxy.internal"),
			},
		},
		{
			name:    "with initial OAuth refresh token",
			options: []ClientOption{WithInitialOAuthRefreshToken("initial"), oauthClientOpt},
			expectedClient: Client{
				endpoint:            defaultEndpoint,
				oAuthClientID:       "id",
				oAuthInitialRefresh: "initial",
				client:              oauthMockClient,
				customHeaders:       map[string]string{"User-Agent": defaultUserAgent},
				epURL:               defaultEndpointURL,
			},
		},
		// We can assume that WithHTTPClient() works since it is used in all above cases.
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
			if err != tc.expectedError {
				t.Fatalf("Client initialization: got error %v, want error %v", err, tc.expectedError)
			}

			if err != nil {
				return
			}

			cmpOpts := cmp.Options{cmp.AllowUnexported(Client{}), cmpopts.IgnoreFields(Client{}, "authProvider")}

			if diff := cmp.Diff(tc.expectedClient, *client, cmpOpts); diff != "" {
				t.Fatalf("Unexpected client: (-want +got)\n%s", diff)
			}
		})
	}
}
