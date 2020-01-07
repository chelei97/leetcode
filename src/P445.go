package main

func main() {
	nums1, nums2 := []int{2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,9}, []int{5,6,4,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,9,9,9,9}
	l1 := CreateList(nums1)
	l2 := CreateList(nums2)
	l1 = addTwoNumbers(l1, l2)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return l1
	}
	first, second := l1, l2
	var res1, res2 []int
	if l1 != nil {
		res1 = append(res1, first.Val)
		for first.Next != nil {
			res1 = append(res1, first.Next.Val)
			first = first.Next
		}
	}
	if l2 != nil {
		res2 = append(res2, second.Val)
		for second.Next != nil {
			res2 = append(res2, second.Next.Val)
			second = second.Next
		}
	}
	dommy := &ListNode{}
	dommy.Next = l1
	if len(res1) < len(res2) {
		res1, res2 = res2, res1
		dommy.Next = l2
	}
	add := make([]int, len(res1) - len(res2))
	res2 = append(add, res2...)
	carry := 0
	for i := len(res1) - 1; i >= 0; i -- {
		tmp := res1[i] + res2[i] + carry
		carry = tmp / 10
		res1[i] = tmp % 10
	}
	if carry > 0 {
		res1 = append([]int{1}, res1...)
	}
	first = dommy.Next
	for i := 0; i < len(res1); i ++ {
		first.Val = res1[i]
		if first.Next == nil && i == len(res1) - 2 {
			tmp := &ListNode{}
			tmp.Val = res1[i + 1]
			first.Next = tmp
			return dommy.Next
		}
		first = first.Next
	}
	return dommy.Next
}
