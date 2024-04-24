package bleemeo

import (
	"context"
	"encoding/json"
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
