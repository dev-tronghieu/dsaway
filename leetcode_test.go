package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{1, 2}
	actual := twoSum(nums, target)
	assert.Equal(t, expected, actual)

	nums = []int{2, 3, 4}
	target = 6
	expected = []int{1, 3}
	actual = twoSum(nums, target)
	assert.Equal(t, expected, actual)

	nums = []int{-1, 0}
	target = -1
	expected = []int{1, 2}
	actual = twoSum(nums, target)
	assert.Equal(t, expected, actual)
}

func twoSum(numbers []int, target int) []int {
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
