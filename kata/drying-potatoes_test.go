package kata

import "testing"

func TestPotatoes(t *testing.T) {
	result := Potatoes(99, 100, 98)
	if result != 50 {
		t.Errorf("Result incorrect; got %d wanted %d", result, 50)
	}
}
