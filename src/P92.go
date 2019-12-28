package main

import(
	"fmt"
)

/**
翻转链表指定位置
 */
func main(){
	nums := []int{1,2,3,4,5}
	head := CreateList(nums)
	head = reverseBetween(head, 2, 4)
	for head != nil {
		fmt.Print(head.Val, "--->")
		head = head.Next
	}
	fmt.Print("\n")
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head.Next == nil || m == n {
		return head
	}
	//返回这个节点
	forreturn := &ListNode{}
	forreturn.Next = head
	pp := forreturn
	//找到反转的起点
	count := 1
	for count < m {
		pp = pp.Next
		head = head.Next
		count ++
	}
	now := head.Next
	pre := head
	next := now.Next
	for count < n {
		next = now.Next
		now.Next = pre
		pre = now
		now = next
		count ++
	}
	head.Next = next
	pp.Next = pre
	return forreturn.Next
}

