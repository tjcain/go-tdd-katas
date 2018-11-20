package kata

// SplitStrings splits the string into pairs of two characters. If the string
// contains an odd number of characters then it should replace the missing
// second character of the final pair with an underscore ('_')
func SplitStrings(input string) []string {
	split := []string{}

	if len(input)%2 != 0 {
		input += "_"
	}

	for i := 0; i < len(input); i += 2 {
		split = append(split, input[i:i+2])
	}
	return split
}
