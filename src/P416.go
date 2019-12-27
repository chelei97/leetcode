package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 5}
	fmt.Println(canPartition(nums))
}

func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	add := 0
	for i := 0; i < len(nums); i ++ {
		add = add + nums[i]
	}
	if add & 1 == 1 {
		return false
	}
	target := add / 2
	dp := make([]bool, target + 1)
	dp[nums[0]] = true
	dp[0] = true
	for i := 1; i < len(nums); i ++ {
		for j := target; j >= 0 && j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j - nums[i]]
		}
		if dp[target] == true {
			return true
		}
	}
	return dp[target]
}
