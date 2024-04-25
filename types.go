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
