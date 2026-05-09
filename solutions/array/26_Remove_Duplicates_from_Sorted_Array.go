package array

/*
https://leetcode.com/problems/remove-duplicates-from-sorted-array?envType=study-plan-v2&envId=top-interview-150
*/
func RemoveDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	x := 1
	//3 2 3 3 4 4 5 6 7
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[x] = nums[i]
			x++
		}
	}
	return x
}
