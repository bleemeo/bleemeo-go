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
	"net/http"
	"net/url"
)

const defaultIteratorPageSize = "2500"

// An Iterator allows browsing all the resources of a specific kind,
// optionally matching specified parameters,
// and automatically fetching the next page when needed.
type Iterator interface {
	// Next sets the iteration cursor on the next resource,
	// fetching the next result page if necessary.
	// It returns whether iteration can continue or not.
	Next(ctx context.Context) bool
	// At returns the resource reached with the last call to Next().
	// If Next() hasn't been called or returned false, At() mustn't be called.
	At() json.RawMessage
	// Err returns the last error that occurred during iteration, if any.
	Err() error
}

func newIterator(c *Client, resource Resource, params url.Values) *iterator {
	if !params.Has("page_size") {
		if params == nil {
			params = url.Values{"page_size": {defaultIteratorPageSize}}
		} else {
			params.Set("page_size", defaultIteratorPageSize)
		}
	}

	return &iterator{
		c:        c,
		resource: resource,
		params:   cloneMap(params),
	}
}

type iterator struct {
	c        *Client
	resource Resource
	params   url.Values

	currentPage  *ResultsPage
	currentIndex int
	err          error
}

func (iter *iterator) Next(ctx context.Context) bool {
	if iter.err != nil {
		return false
	}

	if iter.currentPage == nil || iter.currentIndex >= len(iter.currentPage.Results)-1 {
		if !iter.fetchPage(ctx) {
			return false
		}

		iter.currentIndex = 0
	} else {
		iter.currentIndex++
	}

	return iter.currentIndex < len(iter.currentPage.Results)
}

func (iter *iterator) At() json.RawMessage {
	if iter.currentPage == nil || iter.currentIndex >= len(iter.currentPage.Results) {
		panic("Iterator.At() called in bad conditions")
	}

	return iter.currentPage.Results[iter.currentIndex]
}

func (iter *iterator) Err() error {
	return iter.err
}

func (iter *iterator) fetchPage(ctx context.Context) (ok bool) {
	var (
		reqURI string
		params url.Values
		err    error
	)

	if iter.currentPage == nil { // first fetch
		reqURI, err = url.JoinPath("/", iter.resource)
		if err != nil {
			iter.err = err

			return false
		}

		params = iter.params
	} else {
		if iter.currentPage.Next == "" {
			return false
		}

		reqURI = iter.currentPage.Next
		// Query parameters were given back in the next URL,
		// so no need to re-add them (otherwise, they will grow infinitely).
		params = nil
	}

	_, resp, err := iter.c.Do(ctx, http.MethodGet, reqURI, params, true, nil)
	if err != nil {
		iter.err = err

		return false
	}

	var page ResultsPage

	err = json.Unmarshal(resp, &page)
	if err != nil {
		iter.err = &JSONUnmarshalError{
			jsonError: &jsonError{
				Err:      err,
				DataKind: JsonErrorDataKind_ResultPage,
				Data:     resp,
			},
		}

		return false
	}

	iter.currentPage = &page

	return true
}
