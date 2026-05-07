package array

// Using Kadane's Algorithm

/*
https://leetcode.com/problems/maximum-subarray
*/
func MaxSubArray(nums []int) int {
	maxSum := nums[0]
	currentMax := 0
	for i := 0; i < len(nums); i++ {
		//Approach 1
		//currentMax = max(currentMax, 0)
		//currentMax += nums[i]

		//Approach 2 (faster than above)
		currentMax = max(nums[i], currentMax+nums[i])
		maxSum = max(maxSum, currentMax)
	}
	return maxSum
}
