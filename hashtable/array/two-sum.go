package array

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		r := target - num
		if j, ok := m[r]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return []int{}
}
