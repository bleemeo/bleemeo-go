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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"maps"
	"net/http"
	"net/url"
	"slices"
	"strings"
)

// JSONReaderFrom marshals the given content to JSON,
// and returns a reader to the marshaled data.
func JSONReaderFrom(body any) (io.Reader, error) {
	if body == nil {
		return nil, nil //nolint: nilnil
	}

	data, err := json.Marshal(body)
	if err != nil {
		return nil, &JSONMarshalError{
			jsonError: &jsonError{
				DataKind: JsonErrorDataKind_RequestBody,
				Data:     body,
				Err:      err,
			},
		}
	}

	return bytes.NewReader(data), nil
}

// paramsFromFields builds some [url.Values] from the given fields.
func paramsFromFields(fields []string) url.Values {
	if len(fields) == 0 || len(fields) == 1 && fields[0] == "" {
		return nil
	}

	return url.Values{
		"fields": {strings.Join(fields, ",")},
	}
}

// cleanupResponse ensures we read the whole response to avoid "Connection reset by peer"
// on server, and ensures that the HTTP connection can be reused.
func cleanupResponse(resp *http.Response) {
	_, _ = io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()
}

// unmarshalResponse converts the given body to a [json.RawMessage],
// while ensuring it is actually valid JSON.
// If the given error is not nil, it will be returned immediately
// without any processing, along with a nil [json.RawMessage].
// It takes an int as first parameter just to match the return signature of the [Client.Do] method.
func unmarshalResponse(_ int, respBody []byte, err error) (json.RawMessage, error) {
	if err != nil {
		return nil, err
	}

	raw := make(json.RawMessage, 0, len(respBody))

	err = json.Unmarshal(respBody, &raw)
	if err != nil {
		return nil, &JSONUnmarshalError{
			jsonError: &jsonError{
				Err:      err,
				DataKind: JsonErrorDataKind_RequestBody,
				Data:     respBody[:minInt(len(respBody), errorRespMaxLength)],
			},
		}
	}

	return raw, nil
}

func minInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// cloneMap returns a shallow clone of the given map.
func cloneMap[K comparable, V any](m map[K]V) map[K]V {
	m2 := make(map[K]V, len(m))
	maps.Copy(m2, m)

	return m2
}

// readBodyStart reads the first errorRespMaxLength of the response body.
func readBodyStart(body io.Reader) []byte {
	content, err := io.ReadAll(io.LimitReader(body, errorRespMaxLength))
	if err != nil {
		log.Println("Error reading body:", err)
	}

	return content
}

// makeBadRequestMessage builds a human-readable message from the body of a 400 response.
// The API usually returns a map of field to errors (e.g. {"field": ["error"]}),
// but it may also return a list of errors, a single string or nested structures.
// If the body isn't valid JSON, its (truncated) content is used as-is.
// An empty string is returned when the body holds no information.
func makeBadRequestMessage(body []byte) string {
	var content any

	if err := json.Unmarshal(body, &content); err != nil {
		text := truncatedBody(body)
		if text == "" {
			return ""
		}

		return " " + text
	}

	var sb strings.Builder

	switch value := content.(type) {
	case map[string]any:
		for _, field := range slices.Sorted(maps.Keys(value)) {
			sb.WriteString("\n- ")
			sb.WriteString(field)
			sb.WriteString(": ")
			sb.WriteString(formatBadRequestValue(value[field]))
		}
	case []any:
		for _, item := range value {
			sb.WriteString("\n- ")
			sb.WriteString(formatBadRequestValue(item))
		}
	case nil:
		return ""
	default:
		sb.WriteString(" ")
		sb.WriteString(formatBadRequestValue(value))
	}

	return sb.String()
}

// truncatedBody returns the given response body as a trimmed string,
// truncated to errorBodyMaxLength to be included in an error message.
func truncatedBody(body []byte) string {
	text := strings.TrimSpace(string(body))
	if len(text) > errorBodyMaxLength {
		text = strings.ToValidUTF8(text[:errorBodyMaxLength], "") + "..."
	}

	return text
}

// formatBadRequestValue formats a decoded JSON value from a 400 response on a single line.
func formatBadRequestValue(value any) string {
	switch v := value.(type) {
	case string:
		return v
	case []any:
		items := make([]string, 0, len(v))

		for _, item := range v {
			items = append(items, formatBadRequestValue(item))
		}

		return strings.Join(items, ", ")
	case map[string]any:
		items := make([]string, 0, len(v))

		for _, key := range slices.Sorted(maps.Keys(v)) {
			items = append(items, key+": "+formatBadRequestValue(v[key]))
		}

		return "{" + strings.Join(items, ", ") + "}"
	case nil:
		return "null"
	default:
		return fmt.Sprint(v)
	}
}
