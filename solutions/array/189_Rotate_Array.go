package array

/*
 https://leetcode.com/problems/rotate-array
*/

// Approach 1 (Best) - 3 step reverse (Beats 100%)
func RotateApproach1(nums []int, k int) {
	length := len(nums)
	k = k % length
	reverse(nums, 0, length-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, length-1)
}

func reverse(nums []int, left int, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

// Approach 2 - Cyclic replacement (Beats 22% only)
func RotateApproach2(nums []int, k int) {
	length := len(nums)
	k = k % length
	count := 0 //element replaced so far
	for start := 0; count < length; start++ {
		currentIdx := start
		currentValue := nums[start]
		for {
			nextIdx := (currentIdx + k) % length
			temp := nums[nextIdx]
			nums[nextIdx] = currentValue
			currentValue = temp
			currentIdx = nextIdx
			count++
			if currentIdx == start {
				break //Cycle detected, go for next cycle (start++) if present
			}
		}
	}
}
