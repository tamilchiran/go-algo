package tools

import "sort"

type SortType int

const (
	ASCENDING SortType = iota
	DESCENDING
)

func Sorting(nums []int, sortType SortType) {
	sort.Slice(nums, func(a, b int) bool {
		if sortType == ASCENDING {
			return nums[a] < nums[b]
		}
		return nums[a] > nums[b]
	})
}
