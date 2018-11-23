package kata

import (
	"reflect"
	"testing"
)

func TestPickPeaks(t *testing.T) {
	tt := []struct {
		name  string
		input []int
		want  PosPeaks
	}{
		{
			"1236412321",
			[]int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1},
			PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		},
		{
			"323641232123",
			[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3},
			PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		},
		{
			"32364123212221 - plateau, return first position only",
			[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1},
			PosPeaks{Pos: []int{3, 7, 10}, Peaks: []int{6, 3, 2}},
		},
		{
			"213122221",
			[]int{2, 1, 3, 1, 2, 2, 2, 2, 1},
			PosPeaks{Pos: []int{2, 4}, Peaks: []int{3, 2}},
		},
		{
			"22222132 - should ignore plateaus on the edge of the array",
			[]int{2, 2, 2, 2, 2, 1, 3, 2},
			PosPeaks{Pos: []int{6}, Peaks: []int{3}},
		},
		{
			"21312222 - should ignore plateaus on the edge of the array",
			[]int{2, 1, 3, 1, 2, 2, 2, 2},
			PosPeaks{Pos: []int{2}, Peaks: []int{3}},
		},
		{
			"empty array",
			[]int{},
			PosPeaks{Pos: []int{}, Peaks: []int{}},
		},
		{
			"13,9,-2,-5,8,8,14,-2,-3",
			[]int{13, 9, -2, -5, 8, 8, 14, -2, -3},
			PosPeaks{Pos: []int{6}, Peaks: []int{14}},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := PickPeaks(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("output for %v should be %v; got %v", tc.name, tc.want, got)
			}
		})
	}
}
