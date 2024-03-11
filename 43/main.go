package main

import "fmt"

func main() {
	root := &TreeNode{Key: 1,
		Left: &TreeNode{Key: 2, Left: &TreeNode{Key: 4}, Right: &TreeNode{Key: 5}},
		Right: &TreeNode{Key: 3},
	}

	treeSze:=treeSize(root)

	fmt.Printf("Given tree Size is %d",treeSze)
	fmt.Println()

}

type TreeNode struct {
	Key   int
	Left  *TreeNode
	Right *TreeNode
}

func treeSize(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftSize := treeSize(node.Left)
	rightSize := treeSize(node.Right)
	return leftSize + rightSize + 1
}
