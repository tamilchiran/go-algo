package array

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
