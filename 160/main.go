/* This program demonstrates how to use the reverseBetween function to reverse a sublist of a linked list.
The original list 1 -> 2 -> 3 -> 4 -> 5 \
is reversed from position 2 to position 4, resulting in the list 1 -> 4 -> 3 -> 2 -> 5. */

package main

import "fmt"

func main() {
	// Create a linked list: 1 -> 2 -> 3 -> 4 -> 5
	head := &LinkedList{Val: 1}
	head.Next = &LinkedList{Val: 2}
	head.Next.Next = &LinkedList{Val: 3}
	head.Next.Next.Next = &LinkedList{Val: 4}
	head.Next.Next.Next.Next = &LinkedList{Val: 5}
	fmt.Println("Original list:")
	printList(head)
	m, n := 2, 4
	newHead := reverseLinkedList(head, m, n)

	fmt.Printf("List after reversing from position %d to %d:\n", m, n)
	printList(newHead)

}

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func reverseLinkedList(head *LinkedList, m, n int) *LinkedList {
	if head == nil || m >= n {

		return head
	}
	// Use a dummy node to simplify the reversal process
	dummy := &LinkedList{Val: 0, Next: head}
	prev := dummy

	// Move prev to the node just before the sublist to be reversed
	for i := 0; i < m-1; i++ {
		prev = prev.Next
	}
	// Start is the first node of the sublist to be reversed
	start := prev.Next
	end := start

	// Move end to the node at position n
	for i := 0; i < n-m; i++ {
		end = end.Next

	}

	for start != end {
		temp := start.Next
		start.Next = end.Next
		end.Next = start
		start = temp

	}
	prev.Next = end
	return dummy.Next

}

func printList(head *LinkedList) {
	current := head

	for current != nil {
		fmt.Printf("%d", current.Val)
		current = current.Next
	}
	fmt.Println()

}
