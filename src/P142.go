package main

import "fmt"

func main(){
	root := &ListNode{
		Val:  3,
		Next: nil,
	}
	root.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	root.Next.Next = &ListNode{
		Val:  0,
		Next: nil,
	}
	root.Next.Next.Next = &ListNode{
		Val:  -4,
		Next: root.Next,
	}

	a := detectCycle(root)
	fmt.Println(a.Val)
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for {
		if fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
		}else {
			return nil
		}
		slow = slow.Next
		if fast == slow {
		    break
		}
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
