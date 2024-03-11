package main

import "fmt"

func main() {

	treeRoot := &TreeNode{
		Key: 23,
		Left: &TreeNode{
			Key: 2,
			Right: &TreeNode{
				Key: 30,
				Left: &TreeNode{
					Key:   30,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	PreTraversal(treeRoot)
	fmt.Println("----------------------")
	PostOrder(treeRoot)
	fmt.Println("----------------------")

	InOrderTraversal(treeRoot)
	k := 3

	fmt.Printf("Nodes at distance %d from the root: ", k)

	PrintNodeAtDistance(treeRoot, k)
	fmt.Println()
}

type TreeNode struct {
	Key   int
	Left  *TreeNode
	Right *TreeNode
}

func PreTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Key)
	PreTraversal(root.Left)
	PreTraversal(root.Right)
}

func InOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	InOrderTraversal(root.Left)
	fmt.Println(root.Key)
	InOrderTraversal(root.Right)

}

func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}

	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Println(root.Key)

}

func PrintNodeAtDistance(node *TreeNode, k int) {
	if node == nil {
		return
	}

	if k == 0 {
		fmt.Println(" ", node.Key)
	} else {
		PrintNodeAtDistance(node.Left, k-1)
		PrintNodeAtDistance(node.Right, k-1)
	}

}
