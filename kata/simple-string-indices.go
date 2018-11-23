package kata

import (
	"errors"
)

func SimpleStringIndices(s string, idx uint) (uint, error) {
	n := 0

	if s[idx] != '(' {
		return 0, errors.New("No closing bracket")
	}

	for i, c := range s[idx:] {
		switch c {
		case '(':
			n++
		case ')':
			n--
		}

		if n == 0 {
			return uint(i) + idx, nil
		}

	}
	return 0, errors.New("No closing bracket")
}
