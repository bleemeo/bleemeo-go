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
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func authMockHandler(*http.Request) (int, []byte, error) {
	return http.StatusOK, []byte(
		`{"access_token": "access", "expires_in": 36000, "token_type":` +
			` "Bearer", "scope": "read write", "refresh_token": "refresh"}`,
	), nil
}

func makeMetricMockHandler(t *testing.T, availableResources int) mockHandler {
	t.Helper()

	makeResults := func(page, pageSize int) []json.RawMessage {
		results := make([]json.RawMessage, pageSize)

		for i := range pageSize {
			results[i] = json.RawMessage(fmt.Sprintf(`{"id": %d}`, (page-1)*pageSize+i+1))
		}

		return results
	}

	return func(r *http.Request) (statusCode int, body []byte, err error) {
		currentPage := 1
		pageSize := 0

		q := r.URL.Query()
		if q.Has("page") {
			currentPage, err = strconv.Atoi(q.Get("page"))
			if err != nil {
				return 0, nil, fmt.Errorf("cannot parse current page %q: %w", q.Get("page"), err)
			}
		}

		if q.Has("page_size") {
			pageSize, err = strconv.Atoi(q.Get("page_size"))
			if err != nil {
				return 0, nil, fmt.Errorf("cannot parse page size %q: %w", q.Get("page_size"), err)
			}
		} else {
			t.Fatal("No page size provided (the default one should be present at least)")
		}

		result := ResultsPage{
			Count: availableResources,
		}

		if pageSize > 0 {
			availablePages := availableResources / pageSize
			if currentPage <= availablePages {
				result.Results = makeResults(currentPage, pageSize)

				if currentPage < availablePages {
					result.Next = fmt.Sprintf("%s/v1/metric/?page=%d&page_size=%d", defaultEndpoint, currentPage+1, pageSize)
				}
			}
		}

		data, err := json.Marshal(result)
		if err != nil {
			return http.StatusInternalServerError, nil, err //nolint:wrapcheck
		}

		return http.StatusOK, data, nil
	}
}

func makeClientMockForIteration(
	t *testing.T, metricHandler mockHandler, extraOpts ...ClientOption,
) (client *Client, requestCounter map[string]int) {
	t.Helper()

	requestCounter = make(map[string]int)
	clientMock := &http.Client{
		Transport: &transportMock{
			handlers: map[string]mockHandler{
				tokenPath:     authMockHandler,
				"/v1/metric/": metricHandler,
			},
			counters: requestCounter,
		},
	}

	client, err := NewClient(append([]ClientOption{
		WithCredentials("u", ""),
		WithHTTPClient(clientMock),
	}, extraOpts...)...)
	if err != nil {
		t.Fatal("Failed to init client:", err)
	}

	return client, requestCounter
}

func TestIterator(t *testing.T) {
	t.Parallel()

	t.Run("normal iteration", func(t *testing.T) {
		t.Parallel()

		const totalResources = 15 // the page size is set to 5

		client, requestCounter := makeClientMockForIteration(t, makeMetricMockHandler(t, totalResources))
		iter := client.Iterator(ResourceMetric, url.Values{"page_size": {"5"}})
		objectsCount := 0

		type retObject struct {
			ID int `json:"id"`
		}

		for iter.Next(context.Background()) {
			objectsCount++

			var retOjb retObject

			err := json.Unmarshal(iter.At(), &retOjb)
			if err != nil {
				t.Fatalf("Failed to unmarshal returned object %q: %v", iter.At(), err)
			}

			if retOjb.ID != objectsCount {
				t.Fatalf("Invalid returned object %+v: want ID=%d", retOjb, objectsCount)
			}
		}

		if err := iter.Err(); err != nil {
			t.Fatal("Iterator error:", err)
		}

		if objectsCount != totalResources {
			t.Fatalf("Expected %d objects, got %d", totalResources, objectsCount)
		}

		expectedRequests := map[string]int{
			tokenPath:     1,
			"/v1/metric/": 3,
		}
		if diff := cmp.Diff(expectedRequests, requestCounter); diff != "" {
			t.Fatalf("Unexpected requests (-want +got):\n%s", diff)
		}
	})

	t.Run("empty iteration", func(t *testing.T) {
		t.Parallel()

		client, requestCounter := makeClientMockForIteration(
			t, makeMetricMockHandler(t, 0),
			WithInitialOAuthRefreshToken("refresh"),
		)
		iter := client.Iterator(ResourceMetric, url.Values{})
		objectsCount := 0

		for iter.Next(context.Background()) {
			objectsCount++
		}

		if err := iter.Err(); err != nil {
			t.Fatal("Iterator error:", err)
		}

		if objectsCount != 0 {
			t.Fatalf("Expected %d objects, got %d", 0, objectsCount)
		}

		expectedRequests := map[string]int{
			tokenPath:     1, // only one call, thanks to the initial refresh token
			"/v1/metric/": 1,
		}
		if diff := cmp.Diff(expectedRequests, requestCounter); diff != "" {
			t.Fatalf("Unexpected requests (-want +got):\n%s", diff)
		}
	})

	t.Run("iteration error", func(t *testing.T) {
		t.Parallel()

		handler := func(*http.Request) (statusCode int, body []byte, err error) {
			return http.StatusInternalServerError, nil, nil
		}

		client, requestCounter := makeClientMockForIteration(t, handler)
		iter := client.Iterator(ResourceMetric, url.Values{})
		objectsCount := 0

		for iter.Next(context.Background()) {
			objectsCount++
		}

		if iter.Err() == nil {
			t.Fatal("Expected error '400 - \"400 Bad Request\"'")
		}

		expectedError := &APIError{
			ReqPath:    "/v1/metric/",
			StatusCode: 500,
			Message:    "500 Internal Server Error",
		}
		if diff := cmp.Diff(expectedError, iter.Err(), cmpopts.EquateEmpty()); diff != "" {
			t.Fatalf("Unexpected error (-want +got):\n%s", diff)
		}

		if objectsCount != 0 {
			t.Fatalf("Expected %d objects, got %d", 0, objectsCount)
		}

		expectedRequests := map[string]int{
			tokenPath:     1,
			"/v1/metric/": 1,
		}
		if diff := cmp.Diff(expectedRequests, requestCounter); diff != "" {
			t.Fatalf("Unexpected requests (-want +got):\n%s", diff)
		}
	})

	t.Run("invalid iteration page", func(t *testing.T) {
		t.Parallel()

		invalidJSONPage := []byte(`{"invalid": "page"`)
		handler := func(*http.Request) (statusCode int, body []byte, err error) {
			return http.StatusOK, invalidJSONPage, nil
		}

		client, requestCounter := makeClientMockForIteration(t, handler)
		iter := client.Iterator(ResourceMetric, url.Values{})
		objectsCount := 0

		for iter.Next(context.Background()) {
			objectsCount++
		}

		if iter.Err() == nil {
			t.Fatal("Expected error '400 - \"400 Bad Request\"'")
		}

		expectedError := &JSONUnmarshalError{
			jsonError: &jsonError{
				// Reproducing the expected error, since we can't build it ourselves
				Err:      json.Unmarshal(invalidJSONPage, new(json.RawMessage)),
				DataKind: JsonErrorDataKind_ResultPage,
				Data:     invalidJSONPage,
			},
		}
		cmpOpts := cmp.Options{cmp.AllowUnexported(JSONUnmarshalError{}, json.SyntaxError{}), cmpopts.EquateEmpty()}

		if diff := cmp.Diff(expectedError, iter.Err(), cmpOpts); diff != "" {
			t.Fatalf("Unexpected error (-want +got):\n%s", diff)
		}

		if objectsCount != 0 {
			t.Fatalf("Expected %d objects, got %d", 0, objectsCount)
		}

		expectedRequests := map[string]int{
			tokenPath:     1,
			"/v1/metric/": 1,
		}
		if diff := cmp.Diff(expectedRequests, requestCounter); diff != "" {
			t.Fatalf("Unexpected requests (-want +got):\n%s", diff)
		}
	})

	t.Run("calling At() without Next()", func(t *testing.T) {
		t.Parallel()

		client, err := NewClient(WithInitialOAuthRefreshToken("r"))
		if err != nil {
			t.Fatal("Failed to create client:", err)
		}

		defer func() {
			if r := recover(); r != nil {
				if r != "Iterator.At() called in bad conditions" {
					t.Fatalf("Unexpected panic message: %v", r)
				}
			}
		}()

		client.Iterator(ResourceMetric, url.Values{}).At()

		t.Fatal("Expected Iterator.At() to panic")
	})
}
