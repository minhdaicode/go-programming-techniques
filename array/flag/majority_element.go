package array

// Time complexity: O(n)
// Space complexity: O(n)
func MajorityElement(nums []int) int {
	flag := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == flag {
			count++
		} else if count == 1 {
			flag = nums[i]
		} else {
			count--
		}
	}
	return flag
}
