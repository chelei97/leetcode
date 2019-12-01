package main

import "fmt"

/**
寻找两个有序数组的中位数
 */

func main(){
	nums1, nums2 := []int{}, []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result float64
	var len1, len2, mid, ind1, ind2, next int
	var flag bool
	len1, len2 = len(nums1), len(nums2)
	mid = (len1 + len2) / 2
	if (len1 + len2) % 2 == 0 {
		mid --
	}
	for i := 0; i < len1 + len2; i ++ {
		if ind1 == len1 {
			ind2 ++
			flag = false
		}else if ind2 == len2 {
			ind1 ++
			flag = true
		}else {
			if nums1[ind1] < nums2[ind2] {
				ind1 ++
				flag = true
			}else {
				ind2 ++
				flag = false
			}
		}
		if i == mid {
			break
		}
	}
	if flag {
		result = float64(nums1[ind1 - 1])
		if (len1 + len2) % 2 != 0 {
			return result
		}
		if ind2 == len2 {
			next = nums1[ind1]
		}else if ind1 == len1 {
			next = nums2[ind2]
		}else {
			if nums2[ind2] < nums1[ind1] {
				next = nums2[ind2]
			}else {
				next = nums1[ind1]
			}
		}
	}else {
		result = float64(nums2[ind2 - 1])
		if (len1 + len2) % 2 != 0 {
			return result
		}
		if ind1 == len1 {
			next = nums2[ind2]
		}else if ind2 == len2 {
			next = nums1[ind1]
		}else {
			if nums2[ind2] < nums1[ind1] {
				next = nums2[ind2]
			}else {
				next = nums1[ind1]
			}
		}
	}
	if (len1 + len2) % 2 == 0 {
		return (result + float64(next)) / 2
	}
	return float64(next)
}
