// To find the distinct numbers in a window of a binary tree
package main

import "fmt"

func main() {
	// Example binary tree
	root := &TreeNode{
		Value:3,
		Left: &TreeNode{
			Value:9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Value:20,
			Left: &TreeNode{
				Value:15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Value:7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	window:=[]int{1,2,3}

	result:=distinctNumberInWindow(root,window)
	fmt.Println(result)

}

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func distinctNumberInWindow(root *TreeNode, window []int) []int {
	result := make([]int, 0)
	for _, Value:= range window {
		count := countDistinct(root, make(map[int]bool), Value)
		result = append(result, count)
	}
	return result

}

func countDistinct(node *TreeNode, seen map[int]bool, k int) int {
	if node == nil {
		return len(seen)
	}

	if _, found := seen[node.Value]; !found {
		seen[node.Value] = true
	}

	leftCount := countDistinct(node.Left, seen, k)
	rightCount := countDistinct(node.Right, seen, k)

	if _, found := seen[node.Value-k]; found {
		delete(seen, node.Value-k)
	}
	return leftCount + rightCount

}
