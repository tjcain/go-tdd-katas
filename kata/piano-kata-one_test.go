package kata

import "testing"

func TestBlackOrWhiteKey(t *testing.T) {
	tt := []struct {
		name     string
		stop     int
		keyColor string
	}{
		{"1", 1, "white"},
		{"5", 5, "black"},
		{"12", 12, "black"},
		{"42", 42, "white"},
		{"86", 86, "black"},
		{"88", 88, "white"},
		{"89", 89, "white"},
		{"92", 92, "white"},
		{"100", 100, "black"},
		{"111", 111, "white"},
		{"200", 200, "black"},
		{"2017", 2017, "white"},
		{"526", 526, "black"},
		{"176", 176, "white"},
		{"23", 23, "white"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := BlackOrWhiteKey(tc.stop)
			if s != tc.keyColor {
				t.Errorf("key color for %v should be %v; got %v",
					tc.name,
					tc.keyColor,
					s)
			}
		})
	}
}
