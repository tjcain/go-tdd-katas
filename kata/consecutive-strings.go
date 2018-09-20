package kata

func LongestConsec(strarr []string, k int) string {
	var out string
	var low int
	for k <= len(strarr) {
		var temp string
		for _, s := range strarr[low:k] {
			temp = temp + s
		}
		if len(temp) > len(out) {
			out = temp
		}
		low++
		k++
	}
	return out
}
