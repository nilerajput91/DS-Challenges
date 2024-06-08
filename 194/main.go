package main

import "fmt"

func main() {
	head := &ListNode{Value: 1, Next: &ListNode{Value: 2, Next: &ListNode{Value: 3, Next: &ListNode{Value: 4, Next: &ListNode{Value: 5, Next: nil}}}}}
	fmt.Println("Original List:")
	printList(head)

	k := 2
	head = reverseGroup(head, k)

	fmt.Printf("List after reversing nodes in %d-group:\n", k)
	printList(head)

}

// printList prints the elements of the linked list.
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Value)
		head = head.Next
	}
	fmt.Println("nil")
}

type ListNode struct {
	Value int
	Next  *ListNode
}

func reverseGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head

	}
	// Create a dummy node to handle cases where k = 1
	dummy := &ListNode{
		Next: head,
	}
	prev := dummy

	for head != nil {
		tail := prev

		for i := 0; i < k; i++ {
			// Check if there are enough nodes to reverse

			tail = tail.Next
			if tail == nil {
				// Not enough nodes to reverse, return the original list
				return dummy.Next
			}
		}
		// Reverse the k-group

		nextGroup := tail.Next
		head, tail := reverse(head, tail)
		prev.Next = head
		tail.Next = nextGroup

		// move to the next k-groups

		prev = tail
		head = nextGroup
	}
	return dummy.Next

}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	current := head

	for prev != tail {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return tail, head
}
