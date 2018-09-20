package kata

func BlackOrWhiteKey(keyPressCount int) string {
	keys := []int{0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1}
	m := map[int]string{0: "white", 1: "black"}
	return m[keys[((keyPressCount-1)%88)%len(keys)]]
}
