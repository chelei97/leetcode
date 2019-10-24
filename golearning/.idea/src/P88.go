package main

import "fmt"

/**
合并两个有效数组
 */

func main(){
	nums1 := []int{1}
	nums2 := []int{}
	m, n := 3, 0
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	end1, end2 := false, false
	if m == 0 && n == 0  {
		return
	}
	if m == 0 {
		for i := 0 ; i < n; i ++ {
			nums1[i] = nums2[i]
		}
	}
	if n == 0 {
		return
	}
	for i, flag1, flag2 := 0, m - 1, n - 1; i < m + n; i ++ {
		if end2 ||(!end1 && nums1[flag1] > nums2[flag2]) {
			nums1[m + n - 1 - i] = nums1[flag1]
			flag1 --
			if flag1 == -1 {
				end1 = true
			}
		}else if end1 || (!end2 && nums1[flag1] <= nums2[flag2]){
			nums1[m + n - 1 - i] = nums2[flag2]
			flag2 --
			if flag2 == -1 {
				end2 = true
			}
		}
	}
}