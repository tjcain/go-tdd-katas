package kata

// Interpreter implements a MiniStringF**k interpreter, which is a derivative of
// the Brainf**k. It contains a memory cell which is initialized at 0.
// '+' Increments the memory cell.
// '.' Output the ASCII value of the memory cell.
func Interpreter(code string) string {
	var m int
	var b []byte

	for _, c := range code {
		switch c {
		case '+':
			m++
		case '.':
			b = append(b, byte(m))
		}
	}
	return string(b)
}
