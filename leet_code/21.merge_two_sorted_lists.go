package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 0 ms
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	isNil := func(n *ListNode) bool {
		return n == nil
	}
	switch {
	case isNil(l1) && isNil(l2):
		return nil
	case isNil(l1):
		return l2
	case isNil(l2):
		return l1
	}

	littleNode, bigNode := l1, l2
	if l1.Val > l2.Val {
		bigNode, littleNode = littleNode, bigNode
	}

	return &ListNode{
		Val:  littleNode.Val,
		Next: mergeTwoLists(littleNode.Next, bigNode),
	}
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println("nil")
}

func main() {
	// l1 := &ListNode{
	// 	// Val: 1,
	// 	// Next: &ListNode{
	// 	// 	Val: 2,
	// 	// 	Next: &ListNode{
	// 	// 		Val: 4,
	// 	// 	},
	// 	// },
	// }
	// l2 := &ListNode{
	// 	Val: 0,
	// Next: &ListNode{
	// 	Val: 3,
	// 	Next: &ListNode{
	// 		Val: 4,
	// 	},
	// },
	// }
	printList(mergeTwoLists(nil, nil))
}
