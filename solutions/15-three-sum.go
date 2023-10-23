package solutions

import "sort"

func ThreeSum(nums []int) [][]int {
	l := len(nums)
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := l - 1
		for j < k {
			sum := nums[j] + nums[k] + nums[i]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return result
}
