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
}
