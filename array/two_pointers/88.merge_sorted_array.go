package array

// Time complexity: O(m +n)
// Space complexity: O(1)
func Merge(nums1 []int, m int, nums2 []int, n int) {
	l := m + n - 1
	i := m - 1
	j := n - 1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[l] = nums1[i]
			i--
		} else {
			nums1[l] = nums2[j]
			j--
		}
		l--
	}
}
