package solutions

func ProductExceptSelf(nums []int) []int {
	product := 1
	zeroCount := 0

	for i := range nums {
		if nums[i] == 0 {
			zeroCount++
		} else {
			product *= nums[i]
		}
	}

	if zeroCount == len(nums) {
		return nums
	}

	if zeroCount > 0 {
		for i := range nums {
			if nums[i] == 0 && zeroCount == 1 {
				nums[i] = product
			} else {
				nums[i] = 0
			}
		}
	} else {
		for i := range nums {
			nums[i] = product / nums[i]
		}
	}

	return nums
}
