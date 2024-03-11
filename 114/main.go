// Nth Node from end of linked list

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

	head.Next.Next.Next = &Node{
		Data: 40,
	}

	head.Next.Next.Next.Next = &Node{
		Data: 50,
	}

	n := 4
	printNthNodeFromEnd(head, n)

}

type Node struct {
	Data int
	Next *Node
}

func printNthNodeFromEnd(head *Node, n int) {

	len := 0
	for curr := head; curr != nil; curr = curr.Next {
		len++

	}

	if len < n {
		return
	}
	curr := head
	for i := 0; i < len-n+1; i++ {
		curr = curr.Next

	}
	fmt.Printf("%d\n", curr.Data)

}
