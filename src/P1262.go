package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{4,1,5,3,1}
	fmt.Println(maxSumDivThree(nums))
}

func maxSumDivThree(nums []int) int {
	record := make([]int, len(nums))
	res := 0
	sum := 0
	for i := 0; i < len(nums); i ++ {
		record[i] = nums[i] % 3
		res = res + record[i]
		sum = sum + nums[i]
	}
	if res % 3 == 0 {
		return sum
	}
	if res % 3 == 2 {
		small1 := oneSmall(nums, record, 2)
		small2 := twoSmall(nums, record, 1)
		if small1 == math.MaxInt32 && small2 == math.MaxInt32 {
			return 0
		}
		if small1 > small2 {
			small1 = small2
		}
		return sum - small1
	}
	if res % 3 == 1 {
		small1 := oneSmall(nums, record, 1)
		small2 := twoSmall(nums, record, 2)
		if small1 == math.MaxInt32 && small2 == math.MaxInt32 {
			return 0
		}
		if small1 > small2 {
			small1 = small2
		}
		return sum - small1
	}
	return 0
}

func oneSmall(nums, record []int, target int) int {
	xiao := math.MaxInt32
	bool := 0
	for i := 0; i < len(nums); i ++ {
		if record[i] == target {
			bool ++
			if xiao > nums[i] {
				xiao = nums[i]
			}
		}
	}
	if bool == 0 {
		return math.MaxInt32
	}
	return xiao
}

func twoSmall(nums, record []int, target int) int {
	xiao1, xiao2 := math.MaxInt32, math.MaxInt32
	bool := 0
	for i := 0; i < len(nums); i ++ {
		if record[i] == target {
			bool ++
			if nums[i] < xiao2 {
				xiao2 = nums[i]
				if xiao1 > xiao2 {
					xiao1, xiao2 = xiao2, xiao1
				}
			}
		}
	}
	if bool < 2 {
		return math.MaxInt32
	}
	return xiao1 + xiao2
}

