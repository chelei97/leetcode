package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4,0,5,-5,3,3,0,-4,-5}
	fmt.Println(threeSumClosest(nums, -2))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2] - target
	for i := 0; i < len(nums) - 2; i ++ {
		if i != 0 && nums[i] == nums[i - 1] {
			continue
		}
		tmp := target - nums[i]
		for left, right := i + 1, len(nums) - 1; left < right; {
			thisres := nums[left] + nums[right] - tmp
			if thisres == 0 {
				return target
			}else if thisres < 0 {
				res = min4(res, thisres)
				left ++
			}else {
				res = min4(res, thisres)
				right --
			}
		}
	}
	return target + res
}

func min4(a, b int) int {
	a1, b1 := a, b
	if a < 0 {
		a1 = -a
	}
	if b < 0 {
		b1 = -b
	}
	if a1 < b1 {
		return a
	}
	return b
}
