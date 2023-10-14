package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	assert.Equal(t, true, containsDuplicate(nums))

	nums = []int{1, 2, 3, 4}
	assert.Equal(t, false, containsDuplicate(nums))

	nums = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	assert.Equal(t, true, containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		if m[v] {
			return true
		}
		m[v] = true
	}
	return false
}
