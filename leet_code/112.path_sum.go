package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Pop[T any](arr *[]T) T {
	if len(*arr) == 0 {
		var zero T
		return zero
	}
	lastElem := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	return lastElem
}

// 0 ms
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if (targetSum-root.Val) == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil {
		if hasPathSum(root.Left, targetSum-root.Val) {
			return true
		}
	}
	if root.Right != nil {
		if hasPathSum(root.Right, targetSum-root.Val) {
			return true
		}
	}
	return false
}

func main() {
	// [5,4,8,11,null,13,4,7,2,null,null,null,1]
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  4,
				Left: nil,
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	target := 22
	fmt.Println(hasPathSum(tree, target), "- result")

}
