package main

import "fmt"

func main() {
	groupSizes := []int{3,3,3,3,3,1,3}
	fmt.Println(groupThePeople(groupSizes))
}

func groupThePeople(groupSizes []int) [][]int {
	record := make(map[int][]int)
	var res [][]int
	for i := 0; i < len(groupSizes); i ++ {
		record[groupSizes[i]] = append(record[groupSizes[i]], i)
	}
	for k, v := range record {
		var tmp []int
		for i := 0; i < len(v); i ++ {
			tmp = append(tmp, v[i])
			if (i + 1) % k == 0 {
				res = append(res, tmp)
				tmp = make([]int, 0)
			}
		}
	}
	return res
}
