package solutions

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		complement := target - v
		if _, ok := m[complement]; ok {
			return []int{m[complement], i}
		}
		m[v] = i
	}
	return []int{}
}
