package main

import "fmt"

func main() {
	obstacleGrid := [][]int{
		{0,0,0},
		{0,1,0},
		{0,0,0},
	}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	width, length := len(obstacleGrid), len(obstacleGrid[0])
	record := make([]int, length)
	record[0] = 1
	//第一行
	for i := 1; i < length; i ++ {
		if obstacleGrid[0][i] != 1 {
			record[i] = record[i - 1]
		}
	}
	for i := 1; i < width; i ++ {
		tmp := make([]int, length)
		for j := 0; j < length; j ++ {
			if obstacleGrid[i][j] == 1 {
				tmp[j] = 0
			}else {
				if j == 0 {
					tmp[j] = record[j]
				}else {
					tmp[j] = tmp[j - 1] + record[j]
				}
			}
		}
		record = tmp
	}
	return record[length - 1]
}
