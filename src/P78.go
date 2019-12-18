package main

import "fmt"

func main(){
	nums := []int{1,2,3}
	fmt.Println(subsets(nums))
}

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	helper(nums, 0, &result, []int{})
	return result
}

func helper(nums []int, first int, result *[][]int, combation []int) {
	if first == len(nums){
		x := make([]int, len(combation))
		copy(x, combation)
		*result = append(*result, x)
		return
	}
	for i := first; i <= len(nums); i ++ {
		if i == len(nums) {
			helper(nums, len(nums), result, combation)
			if len(combation) > 0 {
				combation = combation[:len(combation) - 1]
			}
			continue
		}
		combation = append(combation, nums[i])
		helper(nums, i + 1, result, combation)
		combation = combation[:len(combation) - 1]
	}
}
