package bleemeo

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMakeBodyFrom(t *testing.T) {
	t.Parallel()

	type jsonModel struct {
		Key   string `json:"key"`
		Value int    `json:"value"`
	}

	cases := []struct {
		input          any
		expectedOutput Body
		expectedErr    error
	}{
		{
			input:          map[string]int{"foo": 1, "bar": 2},
			expectedOutput: Body{"foo": 1, "bar": 2},
		},
		{
			input: struct {
				F1 string
				F2 int
				F3 float64
			}{F1: "field", F2: 7, F3: 3.14},
			expectedOutput: Body{"F1": "field", "F2": 7, "F3": 3.14},
		},
		{
			input: jsonModel{
				Key:   "SYD",
				Value: 714,
			},
			expectedOutput: Body{"key": "SYD", "value": 714},
		},
		{
			input:       "invalid",
			expectedErr: ErrBodyNotMapOrStruct,
		},
	}

	for _, testCase := range cases {
		tc := testCase

		t.Run(fmt.Sprintf("%T", tc.input), func(t *testing.T) {
			t.Parallel()

			output, err := MakeBodyFrom(tc.input)
			if !errors.Is(err, tc.expectedErr) {
				t.Fatalf("Unexpected error: want %v, got %v", tc.expectedErr, err)
			}

			if diff := cmp.Diff(output, tc.expectedOutput); diff != "" {
				t.Fatalf("Unexpected output (-want, +got):\n%s", diff)
			}
		})
	}
}
