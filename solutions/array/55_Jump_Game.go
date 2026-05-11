package array

/*
https://leetcode.com/problems/jump-game/?envType=study-plan-v2&envId=top-interview-150
*/
func CanJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	/*
	   eg: [3,2,1,0,4], initial farthest = 0
	   0 -> farthest = max(0, 0 + 3) = 3,
	   1 -> farthest = max(3, 1 + 2) = 3,
	   2 -> farthest = max(3, 2 + 1) = 3,
	   3 -> farthest = max(3, 0 + 0) = 3,
	   4 -> farthest = max(3, 4 + 4) = (Can I even stand on this index i? As there is no way to reach here)
	*/

	/*
	   speed -> o(n)
	   space -> o(2) == o(1)
	   beats 100%
	*/
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i > farthest { // This takes care of "Can I even stand on this index i?" 4 > 3, no way to reach 4, the farthest reachable is 3 only
			return false
		}
		farthest = max(farthest, i+nums[i])
		if farthest >= len(nums)-1 {
			return true
		}
	}
	return false
}
