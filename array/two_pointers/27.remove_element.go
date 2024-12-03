package array

// Time complexity: O(n)
// Space complexity: O(n)
func RemoveElement(nums []int, val int) int {
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}
