package main

import "fmt"

func main() {
	root := &TreeNode{
		Val:   10,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Val:   -3,
		Left:  nil,
		Right: nil,
	}
	root.Left.Left = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root.Left.Right = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root.Right.Right = &TreeNode{
		Val:   11,
		Left:  nil,
		Right: nil,
	}
	root.Left.Left.Left = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root.Left.Left.Right = &TreeNode{
		Val:   -2,
		Left:  nil,
		Right: nil,
	}
	root.Left.Right.Right = &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(pathSum(root, 8))
}

func dfs1(root *TreeNode) {
	if root == nil {
		return
	}
	dfs1(root.Left)
	fmt.Println(root.Val)
	dfs1(root.Right)
}
func pathSum(root *TreeNode, sum int) int {
	var result int
	helper1(root, sum, &result)
	return result
}

func helper1(root *TreeNode, sum int, result *int) {
	if root == nil {
		return
	}
	if sum == root.Val {
		*result  = *result + 1
	}
	helper1(root.Left, sum, result)
	helper1(root.Right, sum, result)
	helper1(root.Left, sum - root.Val, result)
	helper1(root.Right, sum - root.Val, result)
}
