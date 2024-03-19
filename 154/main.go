package main

import "fmt"

func main() {
	// Create a complete binary tree
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}

	count := countNodes(root)
	fmt.Println("count is ", count)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)

	if leftHeight == rightHeight {
		return (1 << leftHeight) + countNodes(root.Right)
	} else {
		return (1 << rightHeight) + countNodes(root.Left)
	}
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + getHeight(root.Left)
}
