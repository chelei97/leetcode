package main

import "fmt"

func main() {
	nums := []int{1,2,7,9,981}
	S := 1000000000
	fmt.Println(findTargetSumWays(nums, S))
}

//迭代方法
//func findTargetSumWays(nums []int, S int) int {
//	record := make(map[int]int)
//	record[nums[0]] = 1
//	record[-nums[0]] = 1
//	for i := 1; i < len(nums); i ++ {
//		result := make(map[int]int)
//		for k, v := range record {
//			result[k + nums[i]] = v + result[k + nums[i]]
//			result[k - nums[i]] = v + result[k - nums[i]]
//		}
//		record = result
//		fmt.Println(record)
//	}
//	return record[S]
//}

//dp
func findTargetSumWays(nums []int, S int) int {
	if S > 1000 || S < -1000 {
		return 0
	}
	dp := make([]int, 2001)
	dp[nums[0] + 1000] += 1
	dp[1000 -nums[0]] += 1
	res := nums[0]
	for i := 1; i < len(nums); i ++ {
		res += nums[i]
		tmp := make([]int, 2001)
		for j := 1000 - res; j <= res + 1000; j ++ {
			if j >= nums[i] {
				tmp[j] += dp[j - nums[i]]
			}
			if j + nums[i] < 2001 {
				tmp[j] += dp[j + nums[i]]
			}
		}
		dp = tmp
	}
	return dp[S + 1000]
}