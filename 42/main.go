package main

import (
	"container/list"
	"fmt"
)

func main() {

	root := &TreeNode{Key: 1,
		Left: &TreeNode{Key: 2, Left: &TreeNode{Key: 4}, Right: &TreeNode{Key: 5}},
		Right: &TreeNode{Key: 3},
	}

	levelOrderTreversal(root)
	fmt.Println()

}

type TreeNode struct {
	Key   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderTreversal(root *TreeNode) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		curremtNode := queue.Remove(queue.Front()).(*TreeNode)
		fmt.Printf("%d", curremtNode.Key)

		if curremtNode.Left != nil {
			queue.PushBack(curremtNode.Left)
		}

		if curremtNode.Right != nil {
			queue.PushBack(curremtNode.Right)
		}
	}
}
