package solutions

func TopKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	freqs := make([]int, k)
	result := make([]int, k)
	var minFreqIdx int
	for num, freq := range m {
		for i, v := range freqs {
			if v < freqs[minFreqIdx] {
				minFreqIdx = i
			}
		}

		if freq > freqs[minFreqIdx] {
			freqs[minFreqIdx] = freq
			result[minFreqIdx] = num
		}
	}

	return result
}
