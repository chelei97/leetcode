package main

import "fmt"

/**
环形链表
 */
func main() {
	nums := []int{1,2,3,4}
	head := createList(nums)
	fmt.Println(hasCycle(head))
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	slow = slow.Next
	fast = fast.Next
	if fast.Next == nil {
		return false
	}
	for fast.Next != nil {
		if fast == slow {
			return true
		}
		if fast.Next == nil {
			return false
		}
		fast = fast.Next
		if fast == slow {
			return true
		}
		slow = slow.Next
	}
	return false
}
type ListNode struct {
	Val int
	Next *ListNode
}
