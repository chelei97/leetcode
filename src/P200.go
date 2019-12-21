package main

import "fmt"

func main(){
	grid := [][]byte{
		{'1','1','1','1','0'},
		{'1','1','0','1','0'},
		{'1','1','0','0','0'},
		{'0','0','0','0','0'},
	}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	result := 0
	if len(grid) <= 0 {
		return result
	}
	for i := 0; i < len(grid); i ++ {
		for j := 0; j < len(grid[0]); j ++ {
			if grid[i][j] == '1' {
				result ++
				dfs(i, j, grid)
			}
		}
	}
	return result
}

func dfs(i, j int, gird [][]byte)  {
	if gird[i][j] == '1' {
		gird[i][j] = '0'
	}
	if i - 1 >= 0 && gird[i - 1][j] == '1' {
		dfs(i - 1, j, gird)
	}
	if i + 1 < len(gird) && gird[i + 1][j] == '1' {
		dfs(i + 1, j, gird)
	}
	if j - 1 >= 0 && gird[i][j - 1] == '1' {
		dfs(i, j - 1, gird)
	}
	if j + 1 < len(gird[0]) && gird[i][j + 1] == '1' {
		dfs(i, j + 1, gird)
	}
}
