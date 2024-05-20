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
	"errors"
	"fmt"
	"reflect"
)

const errorRespMaxLength = 1 << 20 // 1MB

var (
	// ErrTokenIsRefreshOnly is returned when trying to request a new token,
	// but credentials haven't been provided.
	ErrTokenIsRefreshOnly = errors.New("the OAuth token can only be refreshed")
	// errTokenHasNoRefresh is returned when the OAuth access token has no associated refresh token.
	errTokenHasNoRefresh = errors.New("the OAuth token has no refresh")
	// ErrNoAuthMeanProvided is returned when the client has no way to retrieve an OAuth token.
	ErrNoAuthMeanProvided = errors.New("no authentication mean provided")
	// ErrTokenRevoke is returned when the logout operation has not been completed successfully.
	ErrTokenRevoke = errors.New("failed to revoke token")
	// ErrResourceNotFound is returned when the resource with the specified ID doesn't exist (HTTP status 404).
	ErrResourceNotFound = errors.New("resource not found")
)

// JSONErrorDataKind indicates the type of data whose conversion failed.
type JSONErrorDataKind int

//nolint: revive,stylecheck,gofmt,gofumpt,goimports
const (
	JsonErrorDataKind_400Details JSONErrorDataKind = iota
	JsonErrorDataKind_401Details
	JsonErrorDataKind_404Details
	JsonErrorDataKind_ResultPage
	JsonErrorDataKind_RequestBody
)

func (kind JSONErrorDataKind) String() string {
	switch kind {
	case JsonErrorDataKind_400Details:
		return "400 details"
	case JsonErrorDataKind_401Details:
		return "401 details"
	case JsonErrorDataKind_404Details:
		return "404 details"
	case JsonErrorDataKind_ResultPage:
		return "result page"
	case JsonErrorDataKind_RequestBody:
		return "request body"
	default:
		return fmt.Sprintf("unknown JsonErrorDataKind %d", kind)
	}
}

// APIError represents an error returned by the Bleemeo API,
// when the StatusCode is in the 4xx or 5xx range.
type APIError struct {
	ReqPath     string
	StatusCode  int
	ContentType string
	Message     string
	Err         error
	// The first MB of the response, if any.
	Response []byte
}

func (apiErr *APIError) Error() string {
	var errStr string

	if apiErr.StatusCode != 0 {
		errStr += fmt.Sprint(apiErr.StatusCode, " - ")
	}

	if apiErr.Message != "" {
		errStr += apiErr.Message
	}

	if apiErr.Err != nil {
		errStr += " (" + apiErr.Err.Error() + ")"
	}

	return errStr
}

func (apiErr *APIError) Unwrap() error {
	return apiErr.Err
}

// An AuthError holds an error due to unspecified or invalid credentials.
type AuthError struct {
	*APIError
	// ErrorCode is RFC 6749's 'error' parameter.
	ErrorCode string
}

func (authErr *AuthError) Error() string {
	if authErr.Err != nil {
		return "authentication error: " + authErr.Err.Error()
	}

	return "authentication error: " + authErr.APIError.Error()
}

func (authErr *AuthError) Unwrap() error {
	return authErr.APIError
}

type jsonError struct {
	Err      error
	DataKind JSONErrorDataKind
	Data     any
}

func (jsonErr *jsonError) Error() string {
	return jsonErr.DataKind.String() + ": " + jsonErr.Err.Error()
}

func (jsonErr *jsonError) Unwrap() error {
	return jsonErr.Err
}

func (jsonErr *jsonError) Is(other error) bool {
	if err := new(jsonError); errors.As(other, &err) {
		if err.DataKind != jsonErr.DataKind || !reflect.DeepEqual(err.Data, jsonErr.Data) {
			return false
		}

		if errors.Is(err.Err, jsonErr.Err) {
			return true
		}

		return reflect.DeepEqual(err.Err, jsonErr.Err)
	}

	return false
}

// JSONMarshalError represents an error that occurred
// during the serialization of data to JSON.
type JSONMarshalError struct {
	*jsonError
}

func (marshalErr *JSONMarshalError) Error() string {
	return "marshalling " + marshalErr.jsonError.Error()
}

// Is returns whether the given error is the same as this.
func (marshalErr *JSONMarshalError) Is(other error) bool {
	if err := new(JSONMarshalError); errors.As(other, &err) {
		return marshalErr.jsonError.Is(err.jsonError)
	}

	return false
}

// JSONUnmarshalError represents an error that occurred
// during the deserialization of data from JSON.
type JSONUnmarshalError struct {
	*jsonError
}

func (unmarshalErr *JSONUnmarshalError) Error() string {
	return "unmarshalling " + unmarshalErr.jsonError.Error()
}

// Is returns whether the given error is the same as this.
func (unmarshalErr *JSONUnmarshalError) Is(other error) bool {
	if err := new(JSONUnmarshalError); errors.As(other, &err) {
		return unmarshalErr.jsonError.Is(err.jsonError)
	}

	return false
}
