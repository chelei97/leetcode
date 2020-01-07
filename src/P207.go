package main

import "fmt"

func main() {
	prerequisites := [][]int{{0,1}, {0,2}, {1,2}}
	fmt.Println(canFinish(3, prerequisites))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	record := make(map[int][]int)
	status := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i ++ {
		record[prerequisites[i][0]] = append(record[prerequisites[i][0]], prerequisites[i][1])
	}
	for k := range record {
		if !dfs2(k, &status, &record) {
			return false
		}
	}
	return true
}

func dfs2(value int, status *[]int, record *map[int][]int) bool {
	if (*status)[value] == 1 {
		return false
	}
	if (*status)[value] == 3 {
		return true
	}
	(*status)[value] = 1
	for i := 0; i < len((*record)[value]); i ++ {
		if !dfs2((*record)[value][i], status, record) {
			return false
		}
	}
	(*status)[value] = 3
	delete(*record, value)
	return true
}