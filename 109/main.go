//Intersection Point of two linked list

package main

import "fmt"

func main() {

	// Example usage
	// Create first linked list: 1 -> 2 -> 3 -> 4 -> 5 -> 6
	headA := &ListNode{Data: 1}
	headA.Next = &ListNode{Data: 2}
	headA.Next.Next = &ListNode{Data: 3}
	headA.Next.Next.Next = &ListNode{Data: 4}
	headA.Next.Next.Next.Next = &ListNode{Data: 5}
	headA.Next.Next.Next.Next.Next = &ListNode{Data: 6}

	// Create second linked list: 9 -> 10 -> 5 -> 6
	headB := &ListNode{Data: 9}
	headB.Next = &ListNode{Data: 10}
	headB.Next.Next = headA.Next.Next.Next.Next // intersect at node 5

	intersectionPoint := intersctionPointOfNode(headA, headB)
	fmt.Printf("instersctionPoint:%v\n", intersectionPoint)

}

type ListNode struct {
	Data int
	Next *ListNode
}

func intersctionPointOfNode(headA, headB *ListNode) *ListNode {
	lenA := len(headA)
	lenB := len(headB)

	// Traverse the longer list by the difference in lengths

	for lenA > lenB {
		headA = headA.Next
		lenA--
	}

	for lenB > lenA {
		headB = headB.Next
		lenB--
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}

		headA = headA.Next
		headB = headB.Next
	}
	return nil

}

func len(head *ListNode) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}
