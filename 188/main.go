package main

import "fmt"

func main() {
	head := &ListNode{
		Value: 1,
		Next: &ListNode{
			Value: 2,
			Next: &ListNode{
				Value: 3,
				Next: &ListNode{
					Value: 4,
					Next: &ListNode{
						Value: 5,
						Next:  nil,
					},
				},
			},
		},
	}
	fmt.Println("original List:")
	printList(head)

	m := 2
	n := 4

	head = reverseBetween(head, m, n)

	fmt.Printf("List After reversing from %d to %d :\n", m, n)
	printList(head)

}

type ListNode struct {
	Value int
	Next  *ListNode
}

func reverseBetween(head *ListNode, m, n int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Value: 0, Next: head}
	prev := dummy

	for i := 0; i < m-1; i++ {
		prev = prev.Next
	}
	// Reverse the sublist from m to n

	current := prev.Next
	next := current.Next

	for i := 0; i < n-m; i++ {
		current.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
		next = current.Next
	}
	return dummy.Next

}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ->", head.Value)
		head = head.Next
	}
	fmt.Println("")
}
