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
	ErrResourceNotFound = errors.New("resource not found")
)

type apiError struct {
	ReqPath     string
	StatusCode  int
	ContentType string
	Message     string
	Err         error
	// The first MB of the response, if any.
	Response []byte
}

func (apiErr *apiError) Error() string {
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

func (apiErr *apiError) Unwrap() error {
	return apiErr.Err
}

type ClientError struct {
	apiError
}

type ServerError struct {
	apiError
}

type AuthError struct {
	ClientError
	ErrorCode string
}

func (authErr *AuthError) Error() string {
	return fmt.Sprintf("authentication error: %s - %s)", authErr.ErrorCode, authErr.Err.Error())
}

type jsonError struct {
	Err      error
	DataKind string
	Data     any
}

func (jsonErr *jsonError) Error() string {
	return jsonErr.DataKind + ": " + jsonErr.Err.Error()
}

func (jsonErr *jsonError) Unwrap() error {
	return jsonErr.Err
}

type JsonMarshalError struct {
	jsonError
}

func (jme *JsonMarshalError) Error() string {
	return "marshalling " + jme.jsonError.Error()
}

type JsonUnmarshalError struct {
	jsonError
}

func (jme *JsonUnmarshalError) Error() string {
	return "unmarshalling " + jme.jsonError.Error()
}
