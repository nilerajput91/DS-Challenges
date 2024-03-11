//Sorted Insert in a Singly Linked List

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

	x := 5
	head=sortedInsert(head, x)
	printList(head)

}

type Node struct {
	Data int
	Next *Node
}

func sortedInsert(head *Node, data int) *Node {
	newNode := &Node{
		Data: data,
	}

	if head == nil || data < head.Data {
		newNode.Next = head
		return newNode
	}

	curr := head

	for curr.Next != nil && data < curr.Next.Data {
		curr = curr.Next

	}

	newNode.Next = curr.Next
	curr.Next = newNode

	return head

}

func printList(head *Node) {

	for head != nil {
		fmt.Printf("%d \n", head.Data)
		head = head.Next
	}
}
