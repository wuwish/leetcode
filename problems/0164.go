package problems

import (
	"sort"
)

func MaximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	max := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] > max {
			max = nums[i+1] - nums[i]
		}
	}
	return max
}
