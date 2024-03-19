package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int // Stores the final result of level order traversal

	queue := []*TreeNode{root} // Queue to process nodes level by level, starting with the root

	if len(queue) > 0 {
		var level []int    // Stores the nodes at the current level
		size := len(queue) // Number of nodes at the current level

		for i := 0; i < size; i++ {
			node := queue[0]  // Dequeue the first node from the queue
			queue = queue[1:] // Number of nodes at the current level

			// Add the node's value to the current level
			level = append(level, root.Val)

			// Enqueue the left and right children of the current node if they exist

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result

}
