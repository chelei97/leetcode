package main

import "fmt"

func main() {
	nums := []int{2,3,1,2,4,3}
	fmt.Println(minSubArrayLen(7, nums))
}

func minSubArrayLen(s int, nums []int) int {
	if nums[0] >= s {
		return 1
	}
	length := len(nums)
	res := length
	count := nums[0] + nums[1]
	for left, right := 0, 1; right < length - 1; {
		for right < length - 1 && count < s {
			right ++
			count += nums[right]
		}
		res = min5(res, right - left + 1)
		for count >= s && left <= right {
			count -= nums[left]
			left ++
		}
		res = min5(res, right - left + 2)
	}
	return res
}

func min5(a, b int) int {
	if a > b {
		return b
	}
	return a
}
