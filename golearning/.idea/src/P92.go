package main

import "fmt"

/**
翻转链表指定位置
 */
func main(){
	nums := []int{1,0,3,4,0,1}
	head := createList(nums)
	reverseBetween(&head)
	for head != nil {
		fmt.Print(head.Val, "--->")
		head = head.Next
	}
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	count := 0
	temp := head
	var pre, first *ListNode
	pre.Val = 0
	pre.Next = head
	for count < n {
		if count < m{
			pre = temp
			temp = temp.Next
			first = pre.Next
		}else{
			linshi := temp.Next
			temp.Next = nil
			linshi.Next = temp
			temp = linshi
		}
	}
}
type ListNode struct {
	Val int
	Next *ListNode
}