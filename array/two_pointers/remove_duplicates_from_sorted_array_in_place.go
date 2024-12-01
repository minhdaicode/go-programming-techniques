package array

// Time complexity: O(n)
// Space complexity: O(1)
func RemoveDuplicatesInPlace(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	idx := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[idx-2] {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}
