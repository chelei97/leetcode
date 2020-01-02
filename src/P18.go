package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1,0,-5,-2,-2,-4,0,1,-2}
	fmt.Println(fourSum(nums, -9))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums) - 3; i ++ {
		if i != 0 && nums[i] == nums[i - 1] {
			continue
		}
		for j := i + 1; j < len(nums) - 2; j ++ {
			if j != i + 1 && nums[j] == nums[j - 1] {
				continue
			}
			tmp := target - nums[i] - nums[j]
			for left, right := j + 1, len(nums) - 1; left < right; {
				if right != len(nums) - 1 && nums[right] == nums[right + 1] {
					right --
					continue
				}
				if left != j + 1 && nums[left] == nums[left - 1] {
					left ++
					continue
				}
				if nums[left] + nums[right] == tmp {
					res := []int{nums[i], nums[j], nums[left], nums[right]}
					result = append(result, res)
					right --
					left ++
				}else if nums[left] + nums[right] < tmp {
					left ++
				}else {
					right --
				}
			}
		}
	}
	return result
}