//merged two tree

package main

import "fmt"

func main() {
	// Example trees
	t1 := &TreeNode{Value: 1}
	t1.Left = &TreeNode{Value: 3}
	t1.Right = &TreeNode{Value: 2}
	t1.Left.Left = &TreeNode{Value: 5}

	t2 := &TreeNode{Value: 2}
	t2.Left = &TreeNode{Value: 1}
	t2.Right = &TreeNode{Value: 3}
	t2.Left.Right = &TreeNode{Value: 4}
	t2.Right.Right = &TreeNode{Value: 7}

	merged := mergedtwoTree(t1, t2)
	fmt.Println("Merged tree:")
	printTree(merged)

}

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func mergedtwoTree(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	merged := &TreeNode{Value: t1.Value + t2.Value}
	merged.Left = mergedtwoTree(t1.Left, t2.Left)
	merged.Right = mergedtwoTree(t1.Right, t2.Right)
	return merged
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Value)
	printTree(root.Left)
	printTree(root.Right)
}
