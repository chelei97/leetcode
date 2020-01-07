package main

import "fmt"

func main() {
	nums := []int{1,2,3,4}
	head := CreateList(nums)
	reorderList(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

//func reorderList(head *ListNode)  {
//	record := make(map[int]*ListNode)
//	for i := 0; head != nil; i ++ {
//		record[i] = head
//		head = head.Next
//	}
//	length := len(record)
//	for i := 0; i < length / 2; i ++ {
//		record[i].Next = record[length - i - 1]
//		record[length - i - 1].Next = record[i + 1]
//	}
//	record[length / 2].Next = nil
//}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head
	//找到最后一个
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	tmp := slow
	slow = slow.Next
	tmp.Next = nil
	//翻转后半部分
	next := slow.Next
	slow.Next = nil
	pre := slow
	slow = next
	for slow != nil {
		next = slow.Next
		slow.Next = pre
		pre = slow
		slow = next
	}
	slow = pre
	fast = head
	for fast != nil {
		if slow.Next == nil && fast.Next != nil{
			tmp = fast.Next
			fast.Next = slow
			slow.Next = tmp
			break
		}
		slow, fast, slow.Next, fast.Next = slow.Next, fast.Next, fast.Next, slow
	}
}
