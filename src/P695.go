package main

import "fmt"

func main(){
	grid := [][]int{
		{0,0,1,0,0,0,0,1,0,0,0,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,1,1,0,1,0,0,0,0,0,0,0,0},
		{0,1,0,0,1,1,0,0,1,0,1,0,0},
		{0,1,0,0,1,1,0,0,1,1,1,0,0},
		{0,0,0,0,0,0,0,0,0,0,1,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,0,0,0,0,0,0,1,1,0,0,0,0},
	}
	fmt.Println(maxAreaOfIsland(grid))
}

func maxAreaOfIsland(grid [][]int) int {
	result := 0
	final := 0
	if len(grid) <= 0 {
		return result
	}
	for i := 0; i < len(grid); i ++ {
		for j := 0; j < len(grid[0]); j ++ {
			if grid[i][j] == 1 {
				result = 0
				dfs1(i, j, grid, &result)
				if result > final {
					final = result
				}
			}
		}
	}
	return final
}

func dfs1(i, j int, gird [][]int, result *int)  {
	if gird[i][j] == 1 {
		*result ++
		gird[i][j] = 0
	}
	if i - 1 >= 0 && gird[i - 1][j] == 1 {
		dfs1(i - 1, j, gird, result)
	}
	if i + 1 < len(gird) && gird[i + 1][j] == 1 {
		dfs1(i + 1, j, gird, result)
	}
	if j - 1 >= 0 && gird[i][j - 1] == 1 {
		dfs1(i, j - 1, gird, result)
	}
	if j + 1 < len(gird[0]) && gird[i][j + 1] == 1 {
		dfs1(i, j + 1, gird, result)
	}
}
