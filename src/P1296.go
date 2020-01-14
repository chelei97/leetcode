package main

import (
	"fmt"
)

func main()  {
	nums := []int{3,3,2,2,1,1}
	fmt.Println(isPossibleDivide(nums, 4))
}

func isPossibleDivide(nums []int, k int) bool {
	length := len(nums)
	if length == 0 || length % k != 0 {
		return false
	}
	rec := make(map[int]int)
	for i := 0; i < length; i ++ {
		rec[nums[i]] ++
	}
	for len(rec) != 0 {
		begin := min7(rec)
		for i := 0; i < k; i ++ {
			rec[begin + i] --
			if rec[begin + i] == 0 {
				delete(rec, begin + i)
			}else if rec[begin + i] < 0 {
				return false
			}
		}
	}
	return true
}

func min7(rec map[int]int) int {
	res := 9999
	for k := range rec {
		if k < res {
			res = k
		}
	}
	return res
}

