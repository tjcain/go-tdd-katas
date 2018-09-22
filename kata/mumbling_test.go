package kata

import "testing"

func TestAccum(t *testing.T) {
	tt := []struct {
		name   string
		input  string
		output string
	}{
		{"abcd", "abcd", "A-Bb-Ccc-Dddd"},
		{"RqaEzty", "RqaEzty", "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"},
		{"cwAt", "cwAt", "C-Ww-Aaa-Tttt"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := Accum(tc.input)
			if s != tc.output {
				t.Errorf("output for %v should be %v; got %v",
					tc.name,
					tc.output,
					s)
			}
		})
	}
}
