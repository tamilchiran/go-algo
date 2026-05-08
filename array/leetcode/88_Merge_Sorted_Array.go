package array

/*
https://leetcode.com/problems/merge-sorted-array?envType=study-plan-v2&envId=top-interview-150
*/
func Merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1            // last valid element in nums1
	j := n - 1            // last valid element in nums2
	mPointer := m + n - 1 // next element index nums1 full array
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[mPointer] = nums1[i]
			i--
		} else {
			nums1[mPointer] = nums2[j]
			j--
		}
		mPointer--
	}

	for j >= 0 {
		nums1[mPointer] = nums2[j]
		j--
		mPointer--
	}
}
