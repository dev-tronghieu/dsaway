package solutions

func ProductExceptSelf(nums []int) []int {
	preProduct := 1
	postProduct := 1

	l := len(nums)
	result := make([]int, l)

	for i := range nums {
		result[i] = 1
	}

	for i := range nums {
		result[i] = preProduct * result[i]
		result[l-1-i] = postProduct * result[l-1-i]
		preProduct *= nums[i]
		postProduct *= nums[l-1-i]
	}

	return result
}
