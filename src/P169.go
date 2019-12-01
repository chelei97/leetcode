package main

import "fmt"

/**
求众数
 */
func main() {
	nums := []int{1,2,3,3,3}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	mymap := make(map[int]int)
	for i := 0; i < len(nums); i ++ {
		mymap[nums[i]] ++
		if mymap[nums[i]] == len(nums) / 2 {
			return nums[i]
		}
	}
	return 0
}