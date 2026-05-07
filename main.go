package main

import (
	"fmt"

	arraystring "github.com/tamilchiran/go-algo/array-string/leetcode"
	array "github.com/tamilchiran/go-algo/array/leetcode"
	mapString "github.com/tamilchiran/go-algo/map-string/leetcode"
)

func main() {
	fmt.Println("\nAlgo 150 Challenge!!")
	//383. Ransom Note
	ransomNote := "a"
	magazine := "ba"
	fmt.Println("383. Ransom Note: " + fmt.Sprint(arraystring.CanConstruct(ransomNote, magazine)))

	//205. Isomorphic Strings
	s := "egg"
	t := "add"
	fmt.Println("205. Isomorphic Strings: " + fmt.Sprint(mapString.IsIsomorphic(s, t)))

	//53. Maximum Subarray
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("53. Maximum Subarray: " + fmt.Sprint(array.MaxSubArray(nums)))

	//88. Merge Sorted Array
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	array.Merge(nums1, m, nums2, n)
	fmt.Println("88. Merge Sorted Array: " + fmt.Sprint(nums1))
}
