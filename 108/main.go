// Segregate even odd nodes of linked list

package main

import "fmt"

func main() {
	// Example usage
	// Create a linked list: 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Data: 1}
	head.Next = &ListNode{Data: 2}
	head.Next.Next = &ListNode{Data: 3}
	head.Next.Next.Next = &ListNode{Data: 4}
	head.Next.Next.Next.Next = &ListNode{Data: 5}

	fmt.Println("Original list:")
	printList(head)

	// Segregate even and odd nodes
	head = segrigateEvenOdd(head)

	fmt.Println("After segregating even and odd nodes:")
	printList(head)

}

type ListNode struct {
	Data int
	Next *ListNode
}

func segrigateEvenOdd(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// Initialize two pointers for even and odd lists

	evenHead := &ListNode{}
	oddHead := &ListNode{}
	evenTail := evenHead
	oddTail := oddHead

	curr := head
	// Traverse the original list and separate even and odd node
	for curr != nil {
		if curr.Data%2 == 0 {
			evenTail.Next = curr
			evenTail = curr
		} else {
			oddTail.Next = curr
			oddTail = curr
		}
		curr = curr.Next
	}

	// Terminate both lists
	evenTail.Next = nil
	oddTail.Next = nil

	// Merge even and odd lists
	evenTail.Next = oddHead.Next

	return evenHead.Next

}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Data)
		head = head.Next
	}
	fmt.Println("")
}
