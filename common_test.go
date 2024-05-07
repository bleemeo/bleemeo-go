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

	return http.ReadResponse(bufio.NewReader(bytes.NewReader(respData)), req)
}
