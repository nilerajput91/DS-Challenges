//reverse linked list

package main

import "fmt"

func main() {
	head := &LinkedList{
		Value: 1,
		Next: &LinkedList{
			Value: 2,
			Next: &LinkedList{
				Value: 3,
				Next: &LinkedList{
					Value: 4,
					Next:  nil,
				},
			},
		},
	}

	newHead:=reverseLinkedList(head)
	printList(newHead)

}

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func reverseLinkedList(head *LinkedList) *LinkedList {
	if head == nil {
		return nil
	}

	curr := head
	var prev *LinkedList = nil

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
func printList(head *LinkedList) {

	curr := head
	for curr != nil {
		fmt.Printf("%d\n", curr.Value)
		curr = curr.Next
	}

}
