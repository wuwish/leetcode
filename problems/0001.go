package problems

/*
https://leetcode-cn.com/problems/two-sum/
*/

func TwoSum(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		second := target - nums[i]
		for j := i + 1; j < length; j++ {
			if nums[j] == second {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
