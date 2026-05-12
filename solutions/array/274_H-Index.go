package array

import "sort"

/*
https://leetcode.com/problems/h-index?envType=study-plan-v2&envId=top-interview-150
*/

func HIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}
	sort.Slice(citations, func(a, b int) bool {
		return citations[a] > citations[b]
	})

	/*
		eg: [3,0,6,1,5]
		Post sorted array: [6,5,3,1,0]
		0 -> citations[0] >= i+1=1, citationCount = 1
		1 -> citations[1] >= i+1=2, citationCount = 2
		2 -> citations[2] >= i+1=3, citationCount = 3
		3 -> citations[3] < i+1=4,  citationCount = 3, so break
	*/

	/*
	   speed -> o(nlogn) due to sorting
	   space -> o(n) due to sorting + o(1) for citationCount
	   beats 100%
	*/
	citationCount := 0
	for i := 0; i < len(citations); i++ {
		if citations[i] >= i+1 {
			citationCount++
		} else {
			break
		}
	}
	return citationCount
}
