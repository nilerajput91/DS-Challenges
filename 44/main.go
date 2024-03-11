package main

import (
	"fmt"
	"math"
)

func main() {

	root := &TreeNode{Key: 1,
		Left:  &TreeNode{Key: 2, Left: &TreeNode{Key: 4}, Right: &TreeNode{Key: 5}},
		Right: &TreeNode{Key: 3},
	}

	maxValue := maxInBinaryTree(root)
	fmt.Printf("Maximum value in the binary tree: %d\n", maxValue)

}

type TreeNode struct {
	Key   int
	Left  *TreeNode
	Right *TreeNode
}

func maxInBinaryTree(node *TreeNode) int {
	if node == nil {
		return math.MinInt
	}

	leftSight := maxInBinaryTree(node.Left)
	rightSight := maxInBinaryTree(node.Right)

	maxValue := max(node.Key, max(leftSight, rightSight))

	return maxValue
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
