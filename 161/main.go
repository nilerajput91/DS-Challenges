// removeNthFromEnd removes the nth node from the end of a linked list and returns the head of the modified list
package main

import "fmt"

func main() {

	head := &LinkedList{Val: 1}
	head.Next = &LinkedList{Val: 2}
	head.Next.Next = &LinkedList{Val: 3}
	head.Next.Next.Next = &LinkedList{Val: 4}
	head.Next.Next.Next.Next = &LinkedList{Val: 5}

	fmt.Println("Original List:")
	printList(head)
	n := 2

	newHead := removeNthNodeFromEnd(head, n)
	fmt.Printf("List after removing the %dnd node from the end:\n", n)

	printList(newHead)

}

type LinkedList struct {
	Val  int
	Next *LinkedList
}

// removeNthFromEnd removes the nth node from the end of a linked list and returns the head of the modified list
func removeNthNodeFromEnd(head *LinkedList, n int) *LinkedList {
	dummy := &LinkedList{Next: head}
	first := dummy
	second := dummy

	for i := 0; i < n+1; i++ {
		first = first.Next
	}

	// Move both pointers together until the first pointer reaches the end
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// Remove the nth node by skipping over it
	second.Next = second.Next.Next

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
