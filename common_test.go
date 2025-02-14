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
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

const httpResponseHeader = "HTTP/1.1 %d %s\r\n\n"

var errMockHandlerNotFound = errors.New("mock handler not found")

type mockHandler func(r *http.Request) (statusCode int, body []byte, err error)

type transportMock struct {
	handlers map[string]mockHandler
	counters map[string]int
}

func (tm *transportMock) RoundTrip(req *http.Request) (*http.Response, error) {
	handler, ok := tm.handlers[req.URL.Path]
	if !ok {
		return nil, fmt.Errorf("%w: %q", errMockHandlerNotFound, req.URL.Path)
	}

	tm.counters[req.URL.Path]++

	statusCode, body, err := handler(req)
	if err != nil {
		return nil, err
	}

	respData := append([]byte(fmt.Sprintf(httpResponseHeader, statusCode, http.StatusText(statusCode))), body...)

	return http.ReadResponse(bufio.NewReader(bytes.NewReader(respData)), req) //nolint:wrapcheck
}
