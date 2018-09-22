package kata

func Decode(roman string) int {
	var output int
	key := map[rune]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}

	var nums []int
	for _, numeral := range roman {
		nums = append(nums, key[numeral])
	}

	for i, num := range nums {
		if i != len(nums)-1 {
			if num != 0 && num < nums[i+1] {
				output += nums[i+1] - num
				nums[i+1] = 0
				continue
			}
		}
		output += num
	}
	return output
}
