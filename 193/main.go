package main

import "fmt"

func main() {
	// Create two linked lists representing numbers
	l1 := &ListNode{Value: 2, Next: &ListNode{Value: 4, Next: &ListNode{Value: 3}}}
	l2 := &ListNode{Value: 5, Next: &ListNode{Value: 6, Next: &ListNode{Value: 4}}}

	result := addTwoNumbers(l1, l2)

	for result != nil {
		fmt.Print(result.Value)
		if result.Next != nil {
			fmt.Print("->")
		}
		result = result.Next
	}

}

type ListNode struct {
	Value int
	Next  *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current, carry := dummyHead, 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Value
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Value
			l2 = l2.Next
		}
		carry = sum / 10
		current.Next = &ListNode{
			Value: sum % 10,
		}
		current = current.Next
	}
	return dummyHead.Next
}
