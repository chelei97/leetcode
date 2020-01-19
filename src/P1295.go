package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{555,901,482,1771}
	fmt.Println(findNumbers(nums))
}

func findNumbers(nums []int) int {
	sort.Ints(nums)
	weishu, chushu, res := 1, 10, 0
	for _, k := range nums {
		if k >= chushu / 10 {
			if k < chushu {
				if weishu & 1 == 0 {
					res ++
				}
				continue
			}
			for k >= chushu {
				chushu = chushu * 10
				weishu ++
			}
			if weishu & 1 == 0 {
				res ++
			}
		}
	}
	return res
}
