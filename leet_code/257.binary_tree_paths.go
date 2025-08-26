package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 0 ms
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	res := []string{}

	var findPath func(node *TreeNode, path string)
	findPath = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		val := strconv.Itoa(node.Val)
		if path == "" {
			path = val
		} else {
			path = path + "->" + val
		}

		if node.Left == nil && node.Right == nil {
			res = append(res, path)
			return
		}

		findPath(node.Left, path)
		findPath(node.Right, path)
	}

	findPath(root, "")
	return res
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
	}

	fmt.Println(binaryTreePaths(tree))
}
