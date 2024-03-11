//Detect loop

package main

import "fmt"

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
	node5.Next = nil

	hasLoop := detectLoop(head)

	if hasLoop {
		fmt.Println("Linked list has loop")
	} else {
		fmt.Println("Linkde lis has no loop ")
	}

}

type ListNode struct {
	Data int
	Next *ListNode
}

func detectLoop(head *ListNode) bool {
	visited := make(map[*ListNode]bool)

	for head != nil {
		if visited[head] {
			return true
		}

		visited[head] = true
		head = head.Next
	}
	return false
}
