package bleemeo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

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

func newIterator(c *Client, resource string, params Params) Iterator {
	return &iterator{
		c:        c,
		resource: resource,
		params:   params,
	}
}

type iterator struct {
	c        *Client
	resource string
	params   Params

	currentPage  *ResultsPage
	currentIndex int
	err          error
}

func (iter *iterator) Next(ctx context.Context) bool {
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
	var reqURI string

	if iter.currentPage == nil { // first fetch
		reqURI = fmt.Sprintf("/%s/", iter.resource)
	} else {
		if iter.currentPage.Next == "" {
			return false
		}

		nextURL, err := url.Parse(iter.currentPage.Next) // TODO: extract cursor/page from the URI, taking it raw is too easy
		if err != nil {
			iter.err = fmt.Errorf("failed to parse next page URL: %w", err)

			return false
		}

		reqURI = nextURL.RequestURI()
	}

	resp, err := iter.c.Do(ctx, http.MethodGet, reqURI, iter.params, true, nil)
	if err != nil {
		iter.err = err

		return false
	}

	var page ResultsPage

	err = json.Unmarshal(resp, &page)
	if err != nil {
		iter.err = fmt.Errorf("failed to parse response: %w", err)

		return false
	}

	iter.currentPage = &page

	return true
}
