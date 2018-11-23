package kata

import "testing"

func TestSimpleStringIndices(t *testing.T) {
	tt := []struct {
		name        string
		inputString string
		inputUint   uint
		want        uint
	}{
		{
			"((1)23(45))(aB)", "((1)23(45))(aB)", uint(0), uint(10),
		},
		{
			"((1)23(45))(aB)", "((1)23(45))(aB)", 1, uint(3),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			want, _ := SimpleStringIndices(tc.inputString, tc.inputUint)
			if want != tc.want {
				t.Errorf("%s returned %d wanted %d", tc.name, want, tc.want)
			}
		})
	}
}

func TestSimpleStringIndicesError(t *testing.T) {
	tt := []struct {
		name        string
		inputString string
		inputUint   uint
		want        uint
	}{
		{
			"((1)23(45))(aB)", "((1)23(45))(aB)", uint(2), uint(10),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := SimpleStringIndices(tc.inputString, tc.inputUint)
			if err == nil {
				t.Errorf("%s should return an error", tc.name)
			}
		})
	}
}
