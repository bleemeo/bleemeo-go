package bleemeo

import "encoding/json"

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
