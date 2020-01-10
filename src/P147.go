package main

import "fmt"

type ListNode1 struct {
	Val int
	Next *ListNode1
}

func CreateList(nums []int) *ListNode1 {
	head := &ListNode1{nums[0],nil}
	temp := head
	for i := 1; i < len(nums); i ++ {
		this := ListNode1{nums[i],nil}
		temp.Next = &this
		temp = temp.Next
	}
	return head
}

func toString(head *ListNode1) {
	for head.Next != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Printf("%d\n", head.Val)
}

func main() {
	nums := []int{4,2,1,3}
	head := CreateList(nums)
	toString(head)
	fmt.Println(head.Val)
	head = insertionSortList(head)
	toString(head)
}

func insertionSortList(head *ListNode1) *ListNode1 {
	if head == nil || head.Next == nil {
		return head
	}
	dommy := &ListNode1{}
	completed, uncompleted := head, head.Next
	dommy.Next = completed
	completed.Next = nil
	for uncompleted != nil {
		tmp := uncompleted.Next
		pre := dommy
		t := dommy.Next
		for t != nil && t.Val < uncompleted.Val {
			pre = pre.Next
			t = t.Next
		}
		pre.Next = uncompleted
		if t == nil {
			uncompleted.Next = nil
			uncompleted = tmp
			continue
		}
		uncompleted.Next = t
		uncompleted = tmp
	}
	return dommy.Next
}
