/* Reverse a linked list in groups of size k

Our Task: Given a Linked List, the task is to Reverse Linked List in Groups of size k

Example:

I/P: 10->20->30->40->50           k=3
O/P: 30->10->20->50->40

I/P: 1->2->3->4->5->6->7->8       k=3
O/P: 3->2->1->6->5->4->8->7

I/P: 1->2->3->4->5->6->7->8       k=5
O/P: 5->4->3->2->1->8->7->6
*/

package main

import "fmt"

func main() {
	node := &ListNode{
		Data: 10,
		Next: &ListNode{
			Data: 20,
			Next: &ListNode{
				Data: 30,
				Next: &ListNode{
					Data: 40,
					Next: &ListNode{
						Data: 50,
					},
				},
			},
		},
	}

	result := reverseKGroups(node, 5)

	for result != nil {
		fmt.Printf(" Reverse List of Node:%d\n", result.Data)
		result = result.Next
	}

}

type ListNode struct {
	Data int
	Next *ListNode
}

func reverseKGroups(head *ListNode, k int) *ListNode {

	if !hasNodes(head, k) {
		return head
	}

	var prev *ListNode
	curr := head
	count := 0

	for curr != nil && count < k {
		curr = curr.Next
		count++
	}

	if count == k {
		prev = reverseList(head, k)
		head.Next = reverseKGroups(curr, k)
	}
	return prev

}

func reverseList(head *ListNode, k int) *ListNode {
	var prev, next *ListNode
	curr := head
	count := 0

	for curr != nil && count < k {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		count++

	}
	return prev

}

func hasNodes(head *ListNode, k int) bool {
	count := 0
	curr := head

	for curr != nil {
		count++
		curr = curr.Next
	}

	return count >= k

}
