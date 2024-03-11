//Pair Sum with given BST

package main

import "fmt"

func main() {
	root := &TreeNode{
		Value: 10,
		Left: &TreeNode{
			Value: 5,
			Right: &TreeNode{
				Value: 20,
			},
			Left: &TreeNode{
				Value: 16,
			},
		},
		Right: &TreeNode{
			Value: 40,
		},
	}
	k := 21

	fmt.Printf("TargetValue is present %v\n", findTarget(root, k))

}

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, targetSum int) bool {
	visited := make(map[int]bool)
	return find(root, targetSum, visited)

}

func find(root *TreeNode, k int, visited map[int]bool) bool {
	if root == nil {
		return false
	}

	if visited[k-root.Value] {
		return true
	}

	visited[root.Value] = true

	return find(root.Left, k, visited) || find(root.Right, k, visited)

}
