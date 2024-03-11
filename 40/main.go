package main

import (
	"fmt"
	"math"
)

func main() {

	treeNode := &TreeNode{
		Key: 20,
		Left: &TreeNode{
			Key:   8,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Key: 30,
			Right: &TreeNode{
				Key:   18,
				Right: nil,
				Left:  nil,
			},
			Left: &TreeNode{
				Key:   35,
				Left:  nil,
				Right: nil,
			},
		},
	}

	isValid := isValidBST(treeNode)

	if isValid {
		fmt.Println("valid BST YEE")
	} else {
		fmt.Println("Not valid BST Sorry")
	}

}

type TreeNode struct {
	Key int

	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, math.MinInt, math.MaxInt)
}

func isValidBSTHelper(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Key < min || root.Key > max {
		return false
	}

	return isValidBSTHelper(root.Left, min, root.Key) && isValidBSTHelper(root.Right, max, root.Key)
}
