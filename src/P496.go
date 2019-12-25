package main

import (
	"container/list"
	"fmt"
)

func main() {
	nums1 := []int{4,1,2}
	nums2 := []int{1,3,4,2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	l := list.New()
	record := make(map[int]int, len(nums2))
	for i := 0; i < len(nums2); i ++ {
		if l.Len() == 0 || l.Back().Value.(int) >= nums2[i] {
			l.PushBack(nums2[i])
			continue
		}
		for l.Len() != 0 && l.Back().Value.(int) < nums2[i] {
			record[l.Back().Value.(int)] = nums2[i]
			l.Remove(l.Back())
		}
		l.PushBack(nums2[i])
	}
	for l.Len() != 0 {
		record[l.Back().Value.(int)] = -1
		l.Remove(l.Back())
	}
	for i := 0; i < len(nums1); i ++ {
		nums1[i] = record[nums1[i]]
	}
	return nums1
}