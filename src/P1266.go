package main

import "fmt"

func main() {
	points := [][]int{
		{3,2},
		{-2,2},
	}
	fmt.Println(minTimeToVisitAllPoints(points))
}

func minTimeToVisitAllPoints(points [][]int) int {
	res := 0
	for i := 1; i < len(points); i ++ {
		x := abs(points[i][0] - points[i - 1][0])
		y := abs(points[i][1] - points[i - 1][1])
		if x > y {
			res = res + x
		}else {
			res = res + y
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

