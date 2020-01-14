package main

import (
	"fmt"
)

func main() {
	nums := []int{1,2,5,9}
	fmt.Println(smallestDivisor(nums, 6))
}

func smallestDivisor(nums []int, threshold int) int {
	left, right := 1, nums[0]
	for i := 1; i < len(nums); i ++ {
		if right < nums[i] {
			right = nums[i]
		}
	}
	for left < right {
		mid := left + (right - left) / 2
		tmp := sums(nums, mid, threshold)
		if tmp > threshold {
			left = mid + 1
		}else if tmp <= threshold {
			right = mid
		}
	}
	return right
}

func sums(nums []int, target, threshold int) int {
	res := 0
	for _, v := range nums {
		res = res + (v - 1) / target + 1
		if res > threshold {
			break
		}
	}
	return res
}
