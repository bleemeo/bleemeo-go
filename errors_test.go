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
	"errors"
	"testing"

	"golang.org/x/oauth2"
)

func TestErrorMessages(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "bad request",
			err:      &APIError{StatusCode: 400, Message: "Bad request:\n- name: Required"},
			expected: "Bad request:\n- name: Required",
		},
		{
			name:     "status only",
			err:      &APIError{StatusCode: 503, Message: "503 Service Unavailable"},
			expected: "503 Service Unavailable",
		},
		{
			name: "status with wrapped error",
			err: &APIError{
				StatusCode: 404,
				Message:    "404 Not Found",
				Err:        ErrResourceNotFound,
			},
			expected: "404 Not Found (" + ErrResourceNotFound.Error() + ")",
		},
		{
			name:     "no message",
			err:      &APIError{StatusCode: 418},
			expected: "418 I'm a teapot",
		},
		{
			name:     "no message and unknown status",
			err:      &APIError{StatusCode: 499},
			expected: "499",
		},
		{
			name:     "no message nor status",
			err:      &APIError{Err: errors.New("some error")}, //nolint: err113
			expected: "some error",
		},
		{
			name: "throttle",
			err: &ThrottleError{
				APIError: &APIError{StatusCode: 429, Message: "Too many requests, need to wait for 30s"},
			},
			expected: "Too many requests, need to wait for 30s",
		},
		{
			name: "auth from API",
			err: &AuthError{
				APIError: &APIError{StatusCode: 401, Message: "Invalid token."},
			},
			expected: "authentication error: 401 Unauthorized: Invalid token.",
		},
		{
			name: "auth without details",
			err: &AuthError{
				APIError: &APIError{StatusCode: 401, Message: "401 Unauthorized"},
			},
			expected: "authentication error: 401 Unauthorized",
		},
		{
			name: "auth from token endpoint",
			err: &AuthError{
				APIError: &APIError{
					StatusCode: 400,
					Message:    invalidCredentialsMsg,
					Err: &oauth2.RetrieveError{
						ErrorCode:        invalidGrantCode,
						ErrorDescription: invalidCredentialsMsg,
					},
				},
				ErrorCode: invalidGrantCode,
			},
			expected: `authentication error: 400 Bad Request: Invalid credentials given. ` +
				`(oauth2: "invalid_grant" "Invalid credentials given.")`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := tc.err.Error(); got != tc.expected {
				t.Fatalf("Expected %q, got %q", tc.expected, got)
			}
		})
	}
}
