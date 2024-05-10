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
	"encoding/json"
	"io"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestJsonReaderFrom(t *testing.T) {
	t.Parallel()

	t.Run("valid JSON", func(t *testing.T) {
		reader, err := jsonReaderFrom(Body{"p1": "v1", "p2": 6.3})
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
		data := Body{"f": func() {}} // unlikely but invalid data

		_, err := jsonReaderFrom(data)
		if err == nil {
			t.Fatal("Expected error, got none")
		}

		expectedErr := &JsonMarshalError{
			jsonError{
				Err:      &json.UnsupportedTypeError{},
				DataKind: JsonErrorDataKind_RequestBody,
				Data:     data,
			},
		}
		if diff := cmp.Diff(err, expectedErr, cmp.AllowUnexported(JsonMarshalError{}), cmpopts.IgnoreInterfaces(struct{ reflect.Type }{}), cmpopts.IgnoreTypes(data["f"])); diff != "" {
			t.Fatalf("Unexpected error (-want +got):\n%s", diff)
		}
	})
}
