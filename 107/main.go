//Detect and remove loop in linked list

package main

func main() {

	head := &ListNode{
		Data: 1,
	}
	node2 := &ListNode{
		Data: 2,
	}

	node3 := &ListNode{
		Data: 3,
	}

	node4 := &ListNode{
		Data: 4,
	}

	node5 := &ListNode{
		Data: 5,
	}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node3

	removeLoopUsingFloyedAlgo(head)

}

type ListNode struct {
	Data int
	Next *ListNode
}

func removeLoopUsingFloyedAlgo(head *ListNode) {

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}

	}

	if slow != fast {
		return
	}

	slow = head

	for slow.Next != fast.Next {
		slow = slow.Next
		fast = fast.Next
	}
	fast.Next = nil

}
