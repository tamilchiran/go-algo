package array

/*
https://leetcode.com/problems/product-of-array-except-self?envType=study-plan-v2&envId=top-interview-150
*/
func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	answers := make([]int, n)
	answers[0] = 1
	for i := 1; i < n; i++ {
		answers[i] = answers[i-1] * nums[i-1]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		answers[i] = answers[i] * suffix
		suffix = suffix * nums[i]
	}
	return answers
}
