package leetcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	nums := []int{3, 0, -2, -1, 1, 2}
	res := threeSum(nums)
	assert.Equal(t, [][]int{{-2, 0, 2}, {-1, 0, 1}, {-2, -1, 3}}, res)

	nums = []int{}
	res = threeSum(nums)
	assert.Equal(t, res, [][]int{})

	nums = []int{0, 0, 0}
	res = threeSum(nums)
	assert.Equal(t, res, [][]int{{0, 0, 0}})
}

func threeSum(nums []int) [][]int {
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
