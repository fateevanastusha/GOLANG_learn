package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 0 ms
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []**TreeNode{&root}
	for len(stack) != 0 {
		pp := stack[len(stack)-1]
		node := *pp

		if node == nil {
			stack = stack[:len(stack)-1]
			continue
		}
		if node.Right != nil {
			stack = append(stack, &node.Right)
		}
		if node.Left != nil {
			stack = append(stack, &node.Left)
		}
		if node.Left == nil && node.Right == nil {
			res = append(res, node.Val)
			*pp = nil
			stack = stack[:len(stack)-1]
			continue
		}
	}
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

	fmt.Println(postorderTraversal(tree))
}
