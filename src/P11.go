package main

import (
	"fmt"
)

/**
盛最多水的容器
 */
func main(){
	height := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	var result int
	for left, right := 0, len(height) - 1; left < right; {
		if ((right - left) * compare(height[left], height[right], -1)) > result {
			result = (right - left) * compare(height[left], height[right], -1)
		}
		if height[right] < height[left] {
			right --
		}else {
				left ++
		}
	}
	return result
}

func compare(a int, b int, flag int) int{
	if a * flag > b * flag {
		return a
	}else {
		return b
	}
}