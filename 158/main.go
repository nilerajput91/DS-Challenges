/* Convert Sorted Array to Binary Search Tree is a Divide and Conquer problem where you're given a sorted array and you need to convert
it into a height-balanced binary search tree (BST). */

package main

import "fmt"

func main() {
	nums := []int{-10, -3, 0, 5, 9}

	root := sortedArrayToBST(nums)
	fmt.Printf("result:%v\n", inOrderTraversal(root))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2

	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}

}

func inOrderTraversal(root *TreeNode) []int {
	var result []int

	if root != nil {
		result = append(result, inOrderTraversal(root.Left)...)
		result = append(result, root.Val)
		result = append(result, inOrderTraversal(root.Right)...)
	}

	return result
}
