//Pairwise swap nodes of linked list

package main

import "fmt"

func main() {

	// Example usage
	// Create a linked list: 1 -> 2 -> 3 -> 4 -> 5 -> 6
	head := &ListNode{data: 1}
	head.Next = &ListNode{data: 2}
	head.Next.Next = &ListNode{data: 3}
	head.Next.Next.Next = &ListNode{data: 4}
	head.Next.Next.Next.Next = &ListNode{data: 5}
	head.Next.Next.Next.Next.Next = &ListNode{data: 6}

	fmt.Println("original linked List:")

	printList(head)

	head = pairWiseSwap(head)

	fmt.Println("after pairwise swap:")

	printList(head)

}

type ListNode struct {
	data int
	Next *ListNode
}

func pairWiseSwap(head *ListNode) *ListNode {
	// If the list is empty or has only one node, no swapping is needed
	if head == nil || head.Next == nil {
		return head
	}

	// Initialize pointers for swapping
	dummy := &ListNode{
		Next: head,
	}

	prev := dummy
	curr := head
	for curr != nil && curr.Next != nil {

		// Node to be swapped
		firstNode := curr
		secondNode := curr.Next
		//swapped
		prev.Next = secondNode
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		//move ponters to next pair

		prev = firstNode

		curr = firstNode.Next

	}

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.data)

		head = head.Next
	}

	fmt.Println()
}
