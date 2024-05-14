package bleemeo

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var errUnreadable = errors.New("unreadable")

func makeClientMockForDo(t *testing.T, handler mockHandler) (c *Client, requestCounter map[string]int) {
	t.Helper()

	requestCounter = make(map[string]int)
	clientMock := &http.Client{
		Transport: &transportMock{
			handlers: map[string]mockHandler{
				"/v1/resource/": handler,
			},
			counters: requestCounter,
		},
	}

	c, err := NewClient(WithHTTPClient(clientMock))
	if err != nil {
		t.Fatal("Failed to initialize client:", err)
	}

	return c, requestCounter
}

type unreadableTransportMock struct {
	status int
	i      int
}

func (utm *unreadableTransportMock) RoundTrip(r *http.Request) (*http.Response, error) {
	return http.ReadResponse(bufio.NewReader(utm), r) //nolint:wrapcheck
}

func (utm *unreadableTransportMock) Read(p []byte) (n int, err error) {
	header := fmt.Sprintf(httpResponseHeader, utm.status, http.StatusText(utm.status))

	if utm.i < len(header) {
		copy(p, header[utm.i:])
	}

	utm.i += len(p)

	if utm.i <= 4096 {
		return len(p), nil
	}

	return 0, errUnreadable
}

func TestClientDo(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name           string
		respStatus     int
		respBody       []byte
		expectedStatus int
		expectedBody   string
		expectedErr    error
	}{
		{
			name:           "ok",
			respStatus:     200,
			respBody:       []byte(`{"id":"1"}`),
			expectedStatus: 200,
			expectedBody:   `{"id":"1"}`,
		},
		{
			name:           "bad request",
			respStatus:     400,
			respBody:       []byte(`{"field": ["Bad usage"]}`),
			expectedStatus: 400,
			expectedErr: &APIError{
				ReqPath:    "/v1/resource/",
				StatusCode: 400,
				Message:    "Bad request:\n- field: Bad usage",
				Response:   []byte(`{"field": ["Bad usage"]}`),
			},
		},
		{
			name:           "unauthorized",
			respStatus:     401,
			respBody:       []byte(`{"error": "invalid_grant", "error_description": "Invalid credentials given."}`),
			expectedStatus: 401,
			expectedErr: &AuthError{
				APIError: &APIError{
					ReqPath:    "/v1/resource/",
					StatusCode: 401,
					Message:    "Invalid credentials given.",
					Response:   []byte(`{"error": "invalid_grant", "error_description": "Invalid credentials given."}`),
				},
				ErrorCode: "invalid_grant",
			},
		},
		{
			name:           "not found",
			respStatus:     404,
			respBody:       []byte(`{"details": "not found"}`),
			expectedStatus: 404,
			expectedErr: &APIError{
				ReqPath:    "/v1/resource/",
				StatusCode: 404,
				Message:    "404 Not Found",
				Err:        fmt.Errorf("%w: /v1/resource/", ErrResourceNotFound),
				Response:   []byte(`{"details": "not found"}`),
			},
		},
		{
			name:           "server error",
			respStatus:     500,
			expectedStatus: 500,
			expectedErr: &APIError{
				ReqPath:    "/v1/resource/",
				StatusCode: 500,
				Message:    "500 Internal Server Error",
			},
		},
	}

	for _, testCase := range cases {
		tc := testCase

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client, _ := makeClientMockForDo(t, func(*http.Request) (int, []byte, error) {
				return tc.respStatus, tc.respBody, nil
			})

			statusCode, body, err := client.Do(context.Background(), "", "/v1/resource/", nil, false, nil)
			if statusCode != tc.expectedStatus {
				t.Fatalf("Expected status code to be %d, got %d", tc.expectedStatus, statusCode)
			}

			if string(body) != tc.expectedBody {
				t.Fatalf("Expected body to be %q, got %q", tc.expectedBody, body)
			}

			if diff := cmp.Diff(err, tc.expectedErr, cmpopts.EquateEmpty(), equateErrorStr("*fmt.wrapError")); diff != "" {
				t.Fatalf("Unexpected error (-want +got):\n%s", diff)
			}
		})
	}

	t.Run("unreadable body", func(t *testing.T) {
		t.Parallel()

		client, err := NewClient(WithHTTPClient(&http.Client{Transport: &unreadableTransportMock{status: 200}}))
		if err != nil {
			t.Fatal("Failed to initialize client:", err)
		}

		statusCode, body, err := client.Do(context.Background(), "", "/v1/resource/", nil, false, nil)
		if statusCode != 200 {
			t.Fatalf("Expected status code to be 200, got %d", statusCode)
		}

		if string(body) != "" {
			t.Fatalf("Expected body to be %q, got %q", "", body)
		}

		expectedErr := &APIError{
			ReqPath:    "/v1/resource/",
			StatusCode: 200,
			Message:    "failed to read response body",
			Err:        errUnreadable,
		}
		if diff := cmp.Diff(err, expectedErr, cmpopts.EquateEmpty(), equateErrorStr("*errors.errorString")); diff != "" {
			t.Fatalf("Unexpected error (-want +got):\n%s", diff)
		}
	})
}

// equateErrorStr considers errors of the given type to be equal
// if their string representations are equal.
func equateErrorStr(errType string) cmp.Option {
	filter := func(a, b error) bool {
		return fmt.Sprintf("%T %T", a, b) == errType+" "+errType
	}
	comparer := func(a, b error) bool {
		return a.Error() == b.Error()
	}

	return cmp.FilterValues(filter, cmp.Comparer(comparer))
}
