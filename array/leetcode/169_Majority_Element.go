package array

func MajorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	k := nums[0]
	vote := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == k {
			vote++
		} else {
			vote--
			if vote <= 0 {
				k = nums[i]
				vote = 1
			}
		}
	}
	return k
}
