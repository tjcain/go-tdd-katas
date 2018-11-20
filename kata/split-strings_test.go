package kata

import (
	"reflect"
	"testing"
)

func TestSplitStrings(t *testing.T) {
	tt := []struct {
		name   string
		input  string
		output []string
	}{
		{"abc", "abc", []string{"ab", "c_"}},
		{"abcdef", "abcdef", []string{"ab", "cd", "ef"}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := SplitStrings(tc.input)
			if !reflect.DeepEqual(got, tc.output) {
				t.Errorf("output for %v should be %v; got %v", tc.name,
					tc.output,
					got)
			}
		})
	}
}
