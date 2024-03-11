// Reverse a linked list iterative
package main

import "fmt"

func main() {

	head := &Node{
		Data: 10,
	}
	head.Next = &Node{
		Data: 20,
	}
	head.Next.Next = &Node{
		Data: 30,
	}

	result := reverseofLinkedList(head)

	for result != nil {
		fmt.Printf("%d\n", result.Data)
		result = result.Next
	}

}

type Node struct {
	Data int
	Next *Node
}

func reverseofLinkedList(head *Node) *Node {

	curr := head
	var prev *Node

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next

	}
	return prev

}
