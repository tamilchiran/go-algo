package array

/*
https://leetcode.com/problems/remove-element?envType=study-plan-v2&envId=top-interview-150
*/
func RemoveElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
