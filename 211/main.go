package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 3,
	}
	root.Left.Left = &TreeNode{
		Val: 4,
	}
	root.Left.Right = &TreeNode{
		Val: 2,
	}
	root.Right.Right = &TreeNode{
		Val: 5,
	}

	windowSize := 3

	inorderTraversal(root, windowSize)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode, windowSize int) {
	var (
		inorder func(node *TreeNode, window []int)
	)

	inorder = func(node *TreeNode, window []int) {
		if node == nil {
			return
		}

		inorder(node.Left, window)
		window = append(window, node.Val)

		if len(window) > windowSize {
			window = window[1:] //remove the oldest element
		}

		fmt.Println(unique(window), "")
		inorder(node.Right, window)
	}

	window := make([]int, 0)
	inorder(root, window)

}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)

	list := []int{}

	for _, key := range intSlice {
		if _, val := keys[key]; !val {
			keys[key] = true
			list = append(list, key)
		}
	}
	return list
}
