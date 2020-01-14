package main

import "fmt"

func main() {
	connections := [][]int{
		{0,1},
		{0,2},
		{2,3},
		{3,4},
	}
	fmt.Println(makeConnected(5,connections))
}

func makeConnected(n int, connections [][]int) int {
	length := len(connections)
	total := n - 1
	if length < n - 1 {
		return -1
	}
	record := make([]int, n)
	for i := 0; i < n; i ++ {
		record[i] = i
	}
	for i := 0; i < len(connections); i ++ {
		root1 := unionsearch(connections[i][0], &record)
		root2 := unionsearch(connections[i][1], &record)
		if root1 != root2 {
			record[root1] = root2
			total --
		}
	}
	return total
}

//查找根结点
func unionsearch(root int, pre *[]int) int {
	var son, tmp int
	son = root
	//我的上级不是掌门
	for root != (*pre)[root] {
		root = (*pre)[root]
	}

	//我就找他的上级，直到掌门出现
	for son != root {
		tmp = (*pre)[son]
		(*pre)[son] = root
		son = tmp
	}
	return root //掌门驾到~~
}
