package main

import "fmt"

func main() {
	nums := []int{1,1,1,3,3}
	head := CreateList(nums)
	head = deleteDuplicates(head)
	for head != nil {
		fmt.Print(head.Val, "--->")
		head = head.Next
	}
	fmt.Print("\n")
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	tmp := head.Val
	flag := false
	for head != nil && head.Next != nil {
		for head.Next != nil && head.Next.Val == tmp {
			head.Next = head.Next.Next
			flag = true
		}
		head = head.Next
		if flag {
			pre.Next = head
		}else {
			pre = pre.Next
		}
		flag = false
		if head == nil {
			break
		}
		tmp = head.Val
	}
	return dummy.Next
}

