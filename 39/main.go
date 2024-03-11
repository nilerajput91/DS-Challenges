package main

import "fmt"

func main() {

	treeNode := &TreeNode{
		key: 1,
		Left: &TreeNode{
			key: 4,
			Right: &TreeNode{
				key: 12,
				Left: &TreeNode{
					key: 50,
					Left: &TreeNode{
						key:  10,
						Left: nil,
					},
				},
			},
		},
	}

	k := 4

	result := kthSmallestEle(treeNode, k)
	fmt.Println("result:", result)

}

type TreeNode struct {
	key   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallestEle(root *TreeNode, key int) int {
	var result int
	count := 0

	var inorderTraversal func(node *TreeNode)
	inorderTraversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorderTraversal(node.Left)

		count++

		if count == key {
			result = node.key
			return
		}

		inorderTraversal(node.Right)
	}

	inorderTraversal(root)
	return result
}
