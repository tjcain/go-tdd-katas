package kata

import "testing"

func TestInAscOrder(t *testing.T) {
	tt := []struct {
		name   string
		input  []int
		output bool
	}{
		{"one", []int{1, 2, 4, 7, 19}, true},
		{"two", []int{1, 2, 3, 4, 5}, true},
		{"three", []int{1, 6, 10, 18, 2, 4, 20}, false},
		{"four", []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, false},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := InAscOrder(tc.input)
			if s != tc.output {
				t.Errorf("output for %v should be %v; got %v",
					tc.name,
					tc.output,
					s)
			}
		})
	}
}
