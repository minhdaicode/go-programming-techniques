package array

// Time complexity: O(n)
// Space complexity: O(1)
func Rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	count := 0

	for i := 0; count < n; i++ {
		cur := i
		prev := nums[i]

		for {
			nextIdx := (cur + k) % n
			nums[nextIdx], prev = prev, nums[nextIdx]
			cur = nextIdx
			count++
			if i == cur {
				break
			}
		}
	}
}
