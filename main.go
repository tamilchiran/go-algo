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

	//27 Remove Element
	nums3 := []int{3, 2, 2, 3}
	val := 3
	fmt.Println("27. Remove Element: " + fmt.Sprint(array.RemoveElement(nums3, val)))

	//26 Remove Duplicates from Sorted Array
	nums4 := []int{1, 1, 2}
	fmt.Println("26. Remove Duplicates from Sorted Array: " + fmt.Sprint(array.RemoveDuplicates(nums4)))

	//80. Remove Duplicates from Sorted Array II
	nums5 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println("80. Remove Duplicates from Sorted Array II: " + fmt.Sprint(array.RemoveDuplicatesII(nums5)))

	//169. Majority Element
	nums6 := []int{1, 1, 1, 2, 2, 2, 3, 1, 1, 1}
	fmt.Println("169. Majority Element: " + fmt.Sprint(array.MajorityElement(nums6)))

	//189. Rotate Array
	// Approach 1
	nums7 := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	array.RotateApproach1(nums7, k)
	fmt.Println("189. Rotate Array Approach 1: " + fmt.Sprint(nums7))

	// Approach 2
	nums8 := []int{1, 2, 3, 4, 5, 6, 7}
	array.RotateApproach2(nums8, k)
	fmt.Println("189. Rotate Array Approach 2: " + fmt.Sprint(nums8))

	//121. Best Time to Buy and Sell Stock
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("121. Best Time to Buy and Sell Stock: " + fmt.Sprint(array.MaxProfit(prices)))
}
