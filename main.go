package main

import (
	"fmt"
	"leetcode/problems"
)


func main() {
	nums1 := []int{0, 0, 0, 0, 0}
	nums2 := []int{-1, 0, 0, 0, 0, 0, 1}
	result := problems.FindMedianSortedArrays(nums1, nums2)
	fmt.Println(result)
}
