package main

import (
	"fmt"
)

func main() {
	graph := [][]int{{1,2}, {3}, {3}, {}}
	fmt.Println(allPathsSourceTarget(graph))
}

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	if n == 1 {
		return [][]int{{0}}
	}
	record := make([]int, n)
	var res [][]int
	dfs5(record, []int{0}, graph, 0, &res)
	return res
}

func dfs5(record, cur []int, graph [][]int, index int, res *[][]int) {
	if index == len(graph) - 1 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	if record[index] == 1 {
		return
	}
	for i := 0; i < len(graph[index]); i ++ {
		cur = append(cur, graph[index][i])
		record[index] = 1
		dfs5(record, cur, graph, graph[index][i], res)
		record[index] = 0
		cur = cur[ : len(cur) - 1]
	}
}
