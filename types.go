package bleemeo

import (
	"bytes"
	"encoding/json"
	"io"
)

type (
	Fields = []string
	Params = map[string]string
	Body   = map[string]any

	ResultsPage struct {
		Count    int               `json:"count"`
		Next     string            `json:"next"`
		Previous string            `json:"previous"`
		Results  []json.RawMessage `json:"results"`
	}
)

func readerFrom(body Body) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}

	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(data), nil
}
