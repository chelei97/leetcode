package main

import "fmt"

func main()  {
	nums := []int{1,2,3,4,5}
	cost := []int{3,4,5,1,2}
	fmt.Println(canCompleteCircuit(nums, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	for i := 0; i < length; i ++ {
		tmp := gas[i]
		for j := i; j < i + length; j ++ {
			now := j
			if j >= length {
				now = j % length
			}
			if tmp < cost[now] {
				break
			}
			if j == i + length - 1 {
				return i
			}
			next := now + 1
			if next >= length {
				next = 0
			}
			tmp = tmp - cost[now] + gas[next]
		}
	}
	return 0
}
