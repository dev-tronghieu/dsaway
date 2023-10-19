package solutions

func LongestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}

	longest := 0
	for num := range m {
		if !m[num-1] {
			temp := num + 1
			for m[temp] {
				delete(m, temp)
				temp++
			}
			consecutiveLen := temp - num
			if consecutiveLen > longest {
				longest = consecutiveLen
			}
		}
	}

	return longest
}
