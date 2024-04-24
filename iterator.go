package bleemeo

import "encoding/json"

type Iterator interface {
	Next() bool
	At() json.RawMessage
	Err() error
}
