package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0}
	fmt.Println(hIndex(nums))
}

func hIndex(citations []int) int {
	sort.Ints(citations)
	length := len(citations)
	//对第i个数字来说，有len - i篇论文至少引用citations[i]次
	for i := length; i >= 1; i -- {
		if citations[length - i] >= i {
			return i
		}
	}
	return 0
}
