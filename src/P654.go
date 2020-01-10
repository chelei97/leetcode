package main

import "fmt"

func main() {
	nums := []int{3,2,1,6,0,5}
	root := constructMaximumBinaryTree(nums)
	fmt.Println(root.Val, root.Left.Val, root.Right.Val)
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return helper5(nums, 0, len(nums) - 1)
}

func helper5(nums []int, begin, end int) *TreeNode {
	if begin > end {
		return nil
	}
	if begin == end {
		root := &TreeNode{
			Val : nums[begin],
		}
		return root
	}
	max := maxIndex(nums, begin, end)
	root := &TreeNode{
		Val : nums[max],
	}
	root.Left = helper5(nums, begin, max - 1)
	root.Right = helper5(nums, max + 1, end)
	return root
}

func maxIndex(nums []int, begin, end int) int {
	res := begin
	for i := begin; i <= end; i ++ {
		if nums[i] > nums[res] {
			res = i
		}
	}
	return res
}
