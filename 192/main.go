package main

import "fmt"

func main() {
	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}
	node4 := &Node{Value: 4}

	node1.Next = node2
	node1.Random = node3
	node2.Next = node3
	node2.Random = node1
	node3.Next = node4
	node3.Random = node2
	node4.Random = node2

	copiedList := copyRandomList(node1)

	fmt.Println("Original List:")
	printList(node1)
	fmt.Println("\n Copied List:")
	printList(copiedList)

}

type Node struct {
	Value  int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	current := head
	// Step 1: Create a copy of each node and insert it after the original node

	for current != nil {
		copy := &Node{Value: current.Value, Next: current.Next}
		current.Next = copy
		current = copy.Next

	}

	// Step 2: Assign random pointers for the copy nodes

	current = head

	for current != nil {
		if current.Random != nil {
			current.Next.Random = current.Random.Next
		}
		current = current.Next.Next
	}

	// Step 3: Separate the original and copy lists
	current = head

	copyHead := head.Next
	copyCurrent := copyHead

	for current != nil {
		current.Next = current.Next.Next
		if copyCurrent.Next != nil {
			copyCurrent.Next = copyCurrent.Next.Next

		}
		current = current.Next
		copyCurrent = copyCurrent.Next
	}
	return copyHead
}

func printList(head *Node) {
	current := head

	for current != nil {
		randomVal := -1

		if current.Random != nil {
			randomVal = current.Random.Value
		}
		fmt.Printf("Value:%d,Random:%d\n", current.Value, randomVal)
		current = current.Next
	}
}
