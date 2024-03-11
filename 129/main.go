//Top View of Binary Tree

package main

import "container/list"

func main() {

}

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// QueueNode represents a node in the queue for level-order traversal
type QueueNode struct {
	Node  *TreeNode
	HDist int // Horizontal distance of the TreeNode from the root
}

func topViewOfBST(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := list.New() // Initialize queue for level-order traversal

	visited := make(map[int]int) // Map to store top view nodes

	// Push the root node into the queue with horizontal distance 0
	queue.PushBack(QueueNode{
		Node:  root,
		HDist: 0,
	})

	for queue.Len() > 0 {
		node := queue.Front().Value.(QueueNode) // Get the front node from the queue
		queue.Remove(queue.Front())             // Remove the front node from the queue

		// If this horizontal distance is not visited before, add it to the result
		if _, ok := visited[node.HDist]; !ok {
			visited[node.HDist] = node.Node.Value // Store the node value in the map
		}

		// Enqueue left child with horizontal distance decreased by 1
		if node.Node.Left != nil {
			queue.PushBack(QueueNode{Node: node.Node.Left, HDist: node.HDist - 1})
		}

		// Enqueue right child with horizontal distance increased by 1
		if node.Node.Right != nil {
			queue.PushBack(QueueNode{Node: node.Node.Right, HDist: node.HDist + 1})
		}

			// Extract the top view nodes from the map and append them to the result slice
		for i := range visited {
			result = append(result, visited[i])
		}

	}
	return result

}
