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
	"bytes"
	"encoding/json"
	"io"
	"log"
)

// jsonReaderFrom marshals the given content to JSON,
// and returns a reader to the marshaled data.
func jsonReaderFrom(body Body) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}

	data, err := json.Marshal(body)
	if err != nil {
		return nil, &JsonMarshalError{
			jsonError: jsonError{
				DataKind: JsonErrorDataKind_RequestBody,
				Data:     body,
				Err:      err,
			},
		}
	}

	return bytes.NewReader(data), nil
}

func unmarshalResponse(_ int, respBody []byte, err error) (json.RawMessage, error) {
	if err != nil {
		return nil, err
	}

	raw := make(json.RawMessage, 0, len(respBody))

	err = json.Unmarshal(respBody, &raw)
	if err != nil {
		return nil, &JsonUnmarshalError{
			jsonError: jsonError{
				Err:      err,
				DataKind: JsonErrorDataKind_RequestBody,
				Data:     respBody[:min(len(respBody), errorRespMaxLength)],
			},
		}
	}

	return raw, nil
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// cloneMap returns a shallow clone of the given map.
func cloneMap[K comparable, V any](m map[K]V) map[K]V {
	m2 := make(map[K]V, len(m))

	for k, v := range m {
		m2[k] = v
	}

	return m2
}

// readBodyStart reads the first errorRespMaxLength of the response body.
func readBodyStart(body io.Reader) []byte {
	content, err := io.ReadAll(io.LimitReader(body, errorRespMaxLength))
	if err == io.EOF {
		err = nil
	} else if err != nil {
		log.Println("Error reading body:", err)
	}

	return content
}
