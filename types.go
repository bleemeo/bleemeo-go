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
	"fmt"
	"reflect"

	"github.com/go-viper/mapstructure/v2"
)

// DefaultFields will make the API to return the model's basic fields.
// Default fields vary from one model to another.
var DefaultFields Fields = nil //nolint:gochecknoglobals, revive

type (
	// Fields represents the list of fields to retrieve for a given model.
	Fields = []string
	// Params represents a set of URL query parameters.
	//
	// For example, a parameter may be a comma-separated field list:
	// [Params]{"fields": "id,label,labels_text"}
	// Note: if no fields are specified, DefaultFields will be used.
	//
	// Or a value to filter a metric listing query:
	// [Params]{"search": "kubernetes"}
	Params = map[string]string
	// A Body represents the data to create or update.
	// It will be marshaled to JSON before being sent to the API.
	Body = map[string]any

	// A ResultsPage represents a section of a resource listing.
	ResultsPage struct {
		// The total number of the requested resource available on the API.
		Count    int               `json:"count"`
		Next     string            `json:"next"`
		Previous string            `json:"previous"`
		Results  []json.RawMessage `json:"results"`
	}
)

// MakeBodyFrom converts the given value to a [Body].
// The value must be a map or a struct.
// Struct fields with json tags are supported.
func MakeBodyFrom(v any) (Body, error) {
	vKind := reflect.ValueOf(v).Kind()
	if vKind != reflect.Map && vKind != reflect.Struct {
		return nil, fmt.Errorf("%w, not a %T", ErrBodyNotMapOrStruct, v)
	}

	var body Body

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{Result: &body, TagName: "json"})
	if err != nil {
		// This error isn't due to the value given by the user;
		// thus, it should never happen in production.
		panic("could not create decoder: " + err.Error())
	}

	err = decoder.Decode(v)
	if err != nil {
		return body, fmt.Errorf("can't convert the given body: %w", err)
	}

	return body, nil
}
