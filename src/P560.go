package main

import "fmt"

func main() {
	nums := []int{1, 1, 1}
	fmt.Println(subarraySum(nums, 2))
}

func subarraySum(nums []int, k int) int {
	res, sum := 0, 0
	record := make(map[int]int)
	record[0] = 1
	for i := 0; i < len(nums); i ++ {
		sum = sum + nums[i]
		if _, ok := record[sum - k]; ok {
			res = res + record[sum - k]
		}
		record[sum]++
	}
	return res
}