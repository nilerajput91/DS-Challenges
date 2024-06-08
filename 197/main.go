//remove duplicate from sorted list

package main

import "fmt"

func main() {

	head := &LinkedList{
		Value: 10,
		Next: &LinkedList{
			Value: 10,
			Next: &LinkedList{
				Value: 20,
				Next: &LinkedList{
					Value: 30,
					Next:  nil,
				},
			},
		},
	}

	sortedList(head)
	printList(head)

}

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func sortedList(head *LinkedList) {

	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Value == curr.Next.Value {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

}

func printList(head *LinkedList) {
	curr := head
	for curr != nil {
		fmt.Printf("%d\n", curr.Value)
		curr = curr.Next
	}
}
