package main

import (
	"container/list"
	"fmt"
)

func main()  {
	nums := []int{1,2,3,4,5}
	fmt.Println(nextGreaterElements(nums))
}

func nextGreaterElements(nums []int) []int {
	l := list.New()
	result := make([]int, len(nums))
	for i := 0; i < 2 * len(nums) - 1; i ++ {
		if l.Len() == 0 || nums[l.Back().Value.(int) % len(nums)] >= nums[i % len(nums)] {
			if l.Len() != 0 && result[l.Back().Value.(int) % len(nums)] != 0 {
				continue
			}
			l.PushBack(i)
			continue
		}
		for l.Len() != 0 && nums[l.Back().Value.(int) % len(nums)] < nums[i % len(nums)] {
			result[l.Back().Value.(int) % len(nums)] = nums[i % len(nums)]
			l.Remove(l.Back())
		}
		l.PushBack(i)
	}
	for l.Len() != 0 {
		if result[l.Back().Value.(int) % len(nums)] == 0 {
			result[l.Back().Value.(int) % len(nums)] = -1
		}
		l.Remove(l.Back())
	}
	return result
}