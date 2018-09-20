package kata

import "testing"

func TestLongestConsec(t *testing.T) {
	tt := []struct {
		name       string
		inputSlice []string
		inputInt   int
		output     string
	}{
		{
			"names",
			[]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"},
			2,
			"wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck",
		},
		{"nill", []string{}, 3, ""},
		{
			"one",
			[]string{"fhiusafgirssgs", "sfhsi", "afghsfghahg"},
			1,
			"fhiusafgirssgs",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := LongestConsec(tc.inputSlice, tc.inputInt)
			if s != tc.output {
				t.Errorf("output for %v should be %v; got %v", tc.name,
					tc.output,
					s)
			}
		})
	}
}
