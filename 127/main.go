// Fix BST with Two Nodes Swapped
package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 2,
			},
		},
	}

	recoverTree(root)

	vals := []int{}

	inorder(root, &vals)
	fmt.Println(vals)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var prev, first, second *TreeNode

	var stack []*TreeNode

	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if prev != nil && curr.Val < prev.Val {
			if first == nil {
				first = prev
			}
			second = curr
		}
		prev = curr
		curr = curr.Right
	}

	first.Val, second.Val = second.Val, first.Val

}

func inorder(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, vals)
	*vals = append(*vals, root.Val)
	inorder(root.Right, vals)

}
