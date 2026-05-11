package array

/*
https://leetcode.com/problems/jump-game-ii/?envType=study-plan-v2&envId=top-interview-150
*/
func JumpII(nums []int) int {

	farthest := 0
	jumps := 0
	currentJumpRangeEnd := 0

	/*
	   eg: [2,3,1,1,4],
	   initial -> farthest = 0,   currentJumpRangeEnd = 0, jumps = 0
	   0 -> farthest = 0 + 2 = 2, currentJumpRangeEnd = 2, jumps = 1
	   1 -> farthest = 1 + 3 = 4, currentJumpRangeEnd = 2, jumps = 1
	   2 -> farthest = 4        , currentJumpRangeEnd = 4, jumps = 2
	   3 -> farthest = 4        , currentJumpRangeEnd = 4, jumps = 2
	   4 -> farthest = 1 + 4 = 5, currentJumpRangeEnd = 5, jumps = 3 // Should not happen, so limit to len(nums)-1
	*/

	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		/*
		   Note: It's guaranteed that you can reach nums[n - 1].
		   If the guarantee doesn't exist, i could go > currentJumpRangeEnd, and can get incorrect jump returned
		*/
		if i == currentJumpRangeEnd { // Can i make another jump within the
			jumps++
			currentJumpRangeEnd = farthest
		}
	}
	return jumps
}
