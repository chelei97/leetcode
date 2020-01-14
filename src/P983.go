package main

import "fmt"

func main() {
	days := []int{2,3,5,6,7,8,9,10,11,17,18,19,23,26,27,29,31,32,33,34,35,36,38,39,40,41,42,43,44,45,47,51,54,55,57,58,64,65,67,68,72,73,74,75,77,78,81,86,87,88,89,91,93,94,95,96,98,99}
	costs := []int{5,24,85}
	fmt.Println(mincostTickets(days, costs))
}

var day []int
var cost []int
var memo [366]int
var record1 map[int]bool
func mincostTickets(days []int, costs []int) int {
	record1 = make(map[int]bool)
	for i := 0; i < len(days); i ++ {
		record1[days[i]] = true
	}
	day = days
	cost = costs
	for i := 0; i < len(memo); i ++ {
		memo[i] = -1
	}
	return dp(1)
}

func dp(i int) int {
	if i > 365 {
		return 0
	}
	if memo[i] != -1 {
		return memo[i]
	}
	if record1[i] == false {
		memo[i] = dp(i + 1)
		return memo[i]
	}
	memo[i] = min9(cost[0] + dp(i + 1), cost[1] + dp(i + 7), cost[2] + dp(i + 30))
	return memo[i]
}

func min9(min1, min2, min3  int) int {
	if min1 > min2 {
		min1 = min2
	}
	if min1 > min3 {
		return min3
	}
	return min1
}








