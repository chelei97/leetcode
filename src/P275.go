package main

import (
	"fmt"
)

func main() {
	nums := []int{0,1}
	fmt.Println(hIndex1(nums))
}

func hIndex1(citations []int) int {
	if len(citations) == 0 || citations[len(citations) - 1] == 0 {
		return 0
	}
	length := len(citations)
	//对第i个数字来说，有len - i篇论文至少引用citations[i]次
	left, right := 1, length
	for left < right {
		mid := left + (right - left) / 2
		if citations[length - mid] < mid {
			right = mid - 1
		}else {
			left = mid + 1
		}
	}
	if citations[length - left] >= left {
		return left
	}
	return left - 1
}
