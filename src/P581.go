package main

import "fmt"

//最短无序连续子数组

func main(){
	nums := []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println(findUnsortedSubarray(nums))
}

func findUnsortedSubarray(nums []int) int {
	begin, end, i, flag := 0, len(nums) - 1, 0, false
	for ; i < len(nums) - 1; i ++ {
		if nums[i] > nums[i + 1]{
			flag = true
			begin = i
			break
		}
	}
	if !flag {
		return 0
	}
	for i = 0; i < len(nums) - 1; i ++ {
		if nums[begin] < nums[i]{
			begin = i
		}
	}
	for i = len(nums) - 1; i > 0; i -- {
		if nums[i] < nums[i - 1] {
			end = i
			break
		}
	}

}