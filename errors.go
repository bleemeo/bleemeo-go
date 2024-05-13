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
)

const errorRespMaxLength = 1 << 20 // 1MB

var (
	ErrTokenIsRefreshOnly = errors.New("the OAuth token can only be refreshed")
	ErrTokenHasNoRefresh  = errors.New("the OAuth token has no refresh")
	ErrTokenRevoke        = errors.New("failed to revoke token")
	ErrResourceNotFound   = errors.New("resource not found")
	ErrBodyNotMapOrStruct = errors.New("the body must be a map or a struct")
)

type JsonErrorDataKind int

const (
	JsonErrorDataKind_401Details JsonErrorDataKind = iota
	JsonErrorDataKind_404Details
	JsonErrorDataKind_ResultPage
	JsonErrorDataKind_RequestBody
)

func (kind JsonErrorDataKind) String() string {
	switch kind {
	case JsonErrorDataKind_401Details:
		return "401 details"
	case JsonErrorDataKind_404Details:
		return "404 details"
	case JsonErrorDataKind_ResultPage:
		return "result page"
	case JsonErrorDataKind_RequestBody:
		return "request body"
	default:
		return fmt.Sprintf("unknown JsonErrorDataKind(%d)", kind)
	}
}

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
		errStr += "\"" + apiErr.Message + "\""
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
	return "authentication error: " + authErr.Err.Error()
}

func (authErr *AuthError) Unwrap() error {
	return authErr.APIError
}

type jsonError struct {
	Err      error
	DataKind JsonErrorDataKind
	Data     any
}

func (jsonErr *jsonError) Error() string {
	return jsonErr.DataKind.String() + ": " + jsonErr.Err.Error()
}

func (jsonErr *jsonError) Unwrap() error {
	return jsonErr.Err
}

type JSONMarshalError struct {
	*jsonError
}

func (jme *JSONMarshalError) Error() string {
	return "marshalling " + jme.jsonError.Error()
}

type JSONUnmarshalError struct {
	*jsonError
}

func (jme *JSONUnmarshalError) Error() string {
	return "unmarshalling " + jme.jsonError.Error()
}
