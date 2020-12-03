package problems

/*
https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
*/

func getMiddle(nums []int, lenNum int) float64 {
	//	偶数长度
	if lenNum%2 == 0 {
		return float64(nums[lenNum/2]+nums[lenNum/2-1]) / 2
	} else {
		return float64(nums[(lenNum-1)/2])
	}
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lenNums1 := len(nums1)
	lenNums2 := len(nums2)
	//	两个数组大小均大于0
	if lenNums1 > 0 && lenNums2 > 0 {
		if nums1[lenNums1-1] <= nums2[0] {
			nums := append(nums1, nums2...)
			return getMiddle(nums, lenNums1+lenNums2)
		} else {
			nums := make([]int, lenNums1+lenNums2)

			i := 0
			j := 0
			z := 0
			for ; i < lenNums1 && j < lenNums2; {
				if nums1[i] <= nums2[j] {
					nums[z] = nums1[i]
					i += 1
				} else {
					nums[z] = nums2[j]
					j += 1
				}
				z += 1
			}
			//	第一个尚未遍历完
			if i < lenNums1 {
				for ; i < lenNums1; i, z = i+1, z+1 {
					nums[z] = nums1[i]
				}
			}
			if j < lenNums2 {
				for ; j < lenNums2; j, z = j+1, z+1 {
					nums[z] = nums2[j]
				}
			}
			return getMiddle(nums, lenNums1+lenNums2)
		}
	} else if lenNums1 > 0 {
		return getMiddle(nums1, lenNums1)
	} else if lenNums2 > 0 {
		return getMiddle(nums2, lenNums2)
	} else {
		return 0
	}
}
