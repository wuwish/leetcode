package problems

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	result := FindMedianSortedArrays(nums1, nums2)
	if result != 2 {
		t.Error("中位数位2")
	}

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	result = FindMedianSortedArrays(nums1, nums2)
	if result != 2.5 {
		t.Error("中位数位2.5")
	}

	nums1 = []int{0, 0, 0, 0, 0}
	nums2 = []int{-1, 0, 0, 0, 0, 0, 1}
	result = FindMedianSortedArrays(nums1, nums2)
	if result != 0 {
		t.Error("中位数位0")
	}
}
