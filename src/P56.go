package main

import (
	"fmt"
	"sort"
)

func main(){
	intervals := [][]int{{1,3},{2,6},{8,10},{15,18}}
	fmt.Println(merge1(intervals))
}

func merge1(intervals [][]int) [][]int {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})
	var result [][]int
	for i := 0; i < len(intervals); {
		j := i + 1
		for j < len(intervals) && intervals[j][0] <= intervals[i][1]{
			intervals[i][1] = Max(intervals[i][1], intervals[j][1])
			j ++
		}
		result = append(result, intervals[i])
		i = j
	}
	return result
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
