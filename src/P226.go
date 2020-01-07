package main

import "fmt"

func main() {
	//count := 1
	root := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	root = invertTree(root)
	fmt.Println(root.Val, root.Left.Val, root.Right.Val)
}

func invertTree(root *TreeNode) *TreeNode {
	helper(root)
	return root
}

func helper(root *TreeNode) {
	if root != nil && root.Left != nil && root.Right != nil {
		root.Right.Val, root.Left.Val = root.Left.Val, root.Right.Val
		helper(root.Right)
		helper(root.Left)
	}
}