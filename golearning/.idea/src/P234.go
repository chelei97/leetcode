package main

import "fmt"

/**
回文链表
 */

func main(){
	nums := []int{1,0,3,4,0,1}
	head := createList(nums)
	fmt.Println(isPalindrome(&head))
}

func createList(nums []int) ListNode {
	head := ListNode{nums[0],nil}
	temp := head
	for i := 1; i < len(nums); i ++ {
		this := ListNode{nums[i],nil}
		temp.Next = &this
		temp = *temp.Next
	}
	return head
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}
	rec := slow.Next
	for rec != nil {
		temp := rec
		rec = temp.Next
		slow.Next = nil
		temp.Next = slow
		slow = temp
	}
	for head != nil && slow != nil {
		if head.Val != slow.Val {
			return false
		}
		head = head.Next
		slow = slow.Next
	}
	return true
}

type ListNode struct {
    Val int
    Next *ListNode
}
