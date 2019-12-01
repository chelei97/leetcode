package main

import (
	"fmt"
	"sort"
)

/**
下一个排列
 */

func main() {
	nums := []int{1,3,2}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int)  {
	length := len(nums)
	m := length - 2
	for ; m >= 0; m -- {
		if nums[m] < nums[m + 1]{
			break
		}
	}
	if m != -1 {
		for j := length - 1; j > m; j -- {
			if nums[j] > nums[m] {
				temp := nums[m]
				nums[m] = nums[j]
				nums[j] = temp
				break
			}
		}
	}
	sort.Ints(nums[m + 1 : ])
}