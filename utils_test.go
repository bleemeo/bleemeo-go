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
	"errors"
	"io"
	"net"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestJsonReaderFrom(t *testing.T) {
	t.Parallel()

	t.Run("valid JSON", func(t *testing.T) {
		t.Parallel()

		reader, err := JSONReaderFrom(map[string]any{"p1": "v1", "p2": 6.3})
		if err != nil {
			t.Fatal("Failed to make reader:", err)
		}

		data, err := io.ReadAll(reader)
		if err != nil {
			t.Fatal("Failed to read from reader:", err)
		}

		expectedData := `{"p1":"v1","p2":6.3}`
		if string(data) != expectedData {
			t.Fatalf("Expected %s but got %s", expectedData, string(data))
		}
	})

	t.Run("invalid JSON", func(t *testing.T) {
		t.Parallel()

		data := map[string]any{"f": func() {}} // unlikely but invalid data

		_, err := JSONReaderFrom(data)
		if err == nil {
			t.Fatal("Expected error, got none")
		}

		expectedErr := &JSONMarshalError{
			&jsonError{
				Err:      &json.UnsupportedTypeError{},
				DataKind: JsonErrorDataKind_RequestBody,
				Data:     data,
			},
		}
		cmpOpts := cmp.Options{
			cmp.AllowUnexported(JSONMarshalError{}),
			cmpopts.IgnoreInterfaces(struct{ reflect.Type }{}),
			cmpopts.IgnoreTypes(func() {}),
		}

		if diff := cmp.Diff(expectedErr, err, cmpOpts); diff != "" {
			t.Fatalf("Unexpected error (-want +got):\n%s", diff)
		}
	})
}

func TestUnmarshalResponse(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name      string
		inputData []byte
		inputErr  error
		// no check required on the expected data
		expectedErr error
	}{
		{
			name:        "valid JSON",
			inputData:   []byte(`{"id": "123"}`),
			inputErr:    nil,
			expectedErr: nil,
		},
		{
			name:      "invalid JSON",
			inputData: []byte(`{"partial":`),
			inputErr:  nil,
			expectedErr: &JSONUnmarshalError{
				jsonError: &jsonError{
					// Reproducing the expected error, since we can't build it ourselves
					Err:      json.Unmarshal([]byte(`{"partial":`), new(json.RawMessage)),
					DataKind: JsonErrorDataKind_RequestBody,
					Data:     []byte(`{"partial":`),
				},
			},
		},
		{
			name:        "response error",
			inputData:   nil,
			inputErr:    net.ErrClosed,
			expectedErr: net.ErrClosed,
		},
	}

	for _, testCase := range cases {
		tc := testCase

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			output, err := unmarshalResponse(0, tc.inputData, tc.inputErr)
			if err == nil && !bytes.Equal(output, tc.inputData) {
				t.Fatalf("Expected the output to be exactly the input, but got %q", string(output))
			}

			if !errors.Is(err, tc.expectedErr) {
				t.Fatalf("Expected the error to be %v, but got %v", tc.expectedErr, err)
			}
		})
	}
}
