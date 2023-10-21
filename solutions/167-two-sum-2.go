package solutions

func TwoSum2(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for l < r {
		sum := numbers[l] + numbers[r]
		if sum < target {
			l++
		} else if sum > target {
			r--
		} else {
			return []int{l + 1, r + 1}
		}
	}

	return nil
}
