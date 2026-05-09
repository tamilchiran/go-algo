package array

func RemoveDuplicatesII(nums []int) int {
	/* My Approach (But not as fast as the one below)
		length := len(nums)
	    if length <= 2 {
	        return length
	    }

	    // Approach 1
	    x := 2
	    //[1,1,1,2,2,3]
	    for i := 2; i < length; i++ {
	        if nums[i] != nums[x-2] {
	            nums[x] = nums[i]
	            x++
	        }
	    }
	    return x
	*/

	// Tried the below best solution from leetcode
	length := len(nums)
	if length == 0 {
		return 0
	}
	i := 1 // read, skip dupes more than 2
	j := 1 // writer
	count := 1
	for i < length {
		if nums[i] == nums[i-1] {
			count++
			if count > 2 {
				i++
				continue
			}
		} else {
			count = 1
		}
		nums[j] = nums[i]
		j++
		i++
	}
	return j
}
