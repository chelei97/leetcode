package main

import "fmt"

func main() {
	nums := []int{1,4,3,2,5,2}
	head := CreateList(nums)
	head = partition(head, 3)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	small, big := &ListNode{}, &ListNode{}
	headsmall, headbig := small, big
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		}else {
			big.Next = head
			big = big.Next
		}
		head = head.Next
	}
	big.Next = nil
	small.Next = nil
	small.Next = headbig.Next
	return headsmall.Next
}