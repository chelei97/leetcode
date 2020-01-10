package main

import "fmt"

func main() {
	nums := []int{3,0,2,1,2}
	fmt.Println(canReach(nums, 2))
}

func canReach(arr []int, start int) bool {
	if len(arr) == 0 {
		return false
	}
	if arr[start] == 0 {
		return true
	}
	rec := make([]int, len(arr))
	var stack []int
	stack = append(stack, start)
	rec[start] = 1
	for len(stack) != 0 {
		var tmp []int
		for i := 0; i < len(stack); i ++ {
			right := stack[i] + arr[stack[i]]
			left := stack[i] - arr[stack[i]]
			if right < len(arr) && right >= 0 && rec[right] == 0 {
				if arr[right] == 0 {
					return true
				}
				tmp = append(tmp, right)
				rec[right] = 1
			}
			if left < len(arr) && left >= 0 && rec[left] == 0 {
				if arr[left] == 0 {
					return true
				}
				tmp = append(tmp, left)
				rec[left] = 1
			}
		}
		stack = tmp
	}
	return false
}
