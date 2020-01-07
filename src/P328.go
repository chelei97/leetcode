package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func CreateList(nums []int) *ListNode {
	head := &ListNode{nums[0],nil}
	temp := head
	for i := 1; i < len(nums); i ++ {
		this := ListNode{nums[i],nil}
		temp.Next = &this
		temp = temp.Next
	}
	return head
}

func main() {
	nums := []int{1,2,3,4}
	head := CreateList(nums)
	oddEvenList(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dommy1 := &ListNode{
		Next: head,
	}
	dommy2 := &ListNode{
		Next: head.Next,
	}
	head = head.Next.Next
	tmp1, tmp2 := dommy1.Next, dommy2.Next
	for head != nil {
		tmp1.Next = head
		tmp1 = tmp1.Next
		head = head.Next
		if head != nil {
			tmp2.Next = head
			tmp2 = tmp2.Next
			head = head.Next
		}else {
			break
		}
	}
	tmp1.Next = dommy2.Next
	tmp2.Next = nil
	return dommy1.Next
}
