package main

import "fmt"

func toString1(this MyLinkedList) {
	head := this.head
	for i := 0; i < this.len - 1; i ++ {
		fmt.Printf("%d -> ", head.val)
		head = head.next
	}
	fmt.Println(head.val)
}

func main() {
	obj := Constructor()
	obj.AddAtHead(4)
	toString1(obj)
	obj.Get(1)
	toString1(obj)
	obj.AddAtHead(1)
	toString1(obj)
	obj.AddAtHead(5)
	toString1(obj)
	obj.DeleteAtIndex(3)
	toString1(obj)
	obj.AddAtHead(7)
	toString1(obj)
	obj.Get(3)
	obj.Get(3)
	obj.Get(3)
	obj.AddAtHead(1)
	obj.DeleteAtIndex(4)
}

type MyLinkedList struct {
	len int
	head *myListNode
}

type myListNode struct {
	val int
	next *myListNode
	pre *myListNode
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	ll := MyLinkedList{
		len : 0,
		head : &myListNode{},
	}
	ll.head.next = ll.head
	ll.head.pre = ll.head
	return ll
}

func (this *MyLinkedList) getNode(index int) *myListNode {
	if index < 0 || index >= this.len {
		return nil
	}
	tmp := this.head
	if index <= this.len / 2 {
		for i := 0; i < index; i ++ {
			tmp = tmp.next
		}
	}else {
		for i := 0; i < this.len - index; i ++ {
			tmp = tmp.pre
		}
	}
	return tmp
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	tmp := this.getNode(index)
	if tmp == nil {
		return -1
	}
	return tmp.val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	if this.len == 0 {
		this.head.val = val
		this.len = 1
		return
	}
	this.AddAtTail(val)
	this.head = this.head.pre
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	if this.len == 0 {
		this.head.val = val
		this.len = 1
		return
	}
	tmp := &myListNode{}
	tmp.val = val
	head := this.head
	end := head.pre
	end.next = tmp
	tmp.pre = end
	tmp.next = head
	head.pre = tmp
	this.len ++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index == this.len {
		this.AddAtTail(val)
	}else if index <= 0 {
		this.AddAtHead(val)
	}else if index >= this.len {
		return
	}else {
		pre := this.getNode(index - 1)
		next := this.getNode(index)
		tmp := &myListNode{}
		tmp.val = val
		pre.next = tmp
		tmp.pre = pre
		tmp.next = next
		next.pre = tmp
		this.len ++
	}
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 || index > this.len - 1 {
		return
	}
	if index == 0 {
		this.head = this.head.next
	}
	tmp := this.getNode(index)
	pre := tmp.pre
	next := tmp.next
	pre.next = next
	next.pre = pre
	this.len --
}
