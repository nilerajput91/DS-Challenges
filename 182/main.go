package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 9}

	serliaized := serliaize(root)
	fmt.Println("Serialzie BST:", serliaized)

	_ = Deseriliaze(serliaized)

	fmt.Println("Deserilazied BST", serliaized)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Serialize converts a BST to a string.

func serliaize(root *TreeNode) string {
	var sab strings.Builder

	serializeHelper(root, &sab)

	return sab.String()
}

// serializeHelper performs a pre-order traversal to serialize the tree.
func serializeHelper(node *TreeNode, sb *strings.Builder) {
	if node == nil {
		sb.WriteString("null,")
		return
	}

	sb.WriteString(strconv.Itoa(node.Val))
	sb.WriteString(",")
	serializeHelper(node.Left, sb)
	serializeHelper(node.Right, sb)
}

// Deserialize converts a string to a BST.

func Deseriliaze(str string) *TreeNode {
	node := strings.Split(str, ",")
	return deseriliazeHelper(&node)
}

// deserializeHelper reconstructs the tree from the serialized string.

func deseriliazeHelper(nodes *[]string) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}

	// Get the next node value from the list.
	val := (*nodes)[0]
	// Remove the processed node value.
	*nodes = (*nodes)[1:]

	if val == "null" {
		return nil
	}

	// Convert the node value to an integer.
	intVal, _ := strconv.Atoi(val)
	// Create a new node with the integer value.
	node := &TreeNode{Val: intVal}
	// Recur to build the left and right subtrees.

	node.Left = deseriliazeHelper(nodes)
	node.Right = deseriliazeHelper(nodes)
	return node

}
