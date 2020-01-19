package main

import "fmt"

func main() {
	grid := [][]int{{3,0,8,4},{2,4,5,7},{9,2,6,3},{0,3,1,0}}
	fmt.Println(maxIncreaseKeepingSkyline(grid))
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	max := 0
	res := 0
	maxcol, maxrow := make([]int, len(grid[0])), make([]int, len(grid))
	for i := 0; i < len(grid); i ++ {
		max = grid[i][0]
		for j := 0; j < len(grid[0]); j ++ {
			if max < grid[i][j] {
				max = grid[i][j]
			}
		}
		maxrow[i] = max
	}
	for i := 0; i < len(grid[0]); i ++ {
		max = grid[0][i]
		for j := 0; j < len(grid); j ++ {
			if max < grid[j][i] {
				max = grid[j][i]
			}
		}
		maxcol[i] = max
	}
	for i := 0; i < len(grid); i ++ {
		for j := 0; j < len(grid[0]); j ++ {
			res = res + min10(maxrow[i],maxcol[j]) - grid[i][j]
		}
	}
	return res
}

func min10(a, b int) int  {
	if a < b {
		return a
	}
	return b
}
