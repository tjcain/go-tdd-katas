package kata

import "testing"

func TestDecode(t *testing.T) {
	tt := []struct {
		name     string
		numerals string
		value    int
	}{
		{"XXI", "XXI", 21},
		{"IV", "IV", 4},
		{"MCMXL", "MCMXL", 1940},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := Decode(tc.numerals)
			if s != tc.value {
				t.Errorf("output for %v should be %v; got %v",
					tc.name,
					tc.value,
					s)
			}
		})
	}
}
