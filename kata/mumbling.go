package kata

import "strings"

func Accum(s string) string {
	buffer := strings.Split(s, "")
	for i, c := range buffer {
		buffer[i] = strings.Title(strings.ToLower(strings.Repeat(c, i+1)))
	}
	return strings.Join(buffer, "-")
}
