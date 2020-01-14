package main

import "fmt"

func main() {
	nums := []int{1,2,3,1}
	fmt.Println(containsNearbyAlmostDuplicate(nums, 3, 0))
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	length := len(nums)
	if k > length - 1 {
		k = length - 1
	}
	dp := make([][]int, 0)
	for i := 0; i < length; i ++ {
		tmp := make([]int, k + 1)
		dp = append(dp, tmp)
	}
	for i := 0; i < len(dp); i ++ {
		for j := 0; j < len(dp[0]); j ++ {

		}
	}
}


