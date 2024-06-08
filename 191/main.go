package main

import "fmt"

func main() {
	l1 := &LinkedList{Value: 1, Next: &LinkedList{Value: 2, Next: &LinkedList{Value: 4}}}
	l2 := &LinkedList{Value: 1, Next: &LinkedList{Value: 3, Next: &LinkedList{Value: 4}}}

	merged := merageTwoLinkedList(l1, l2)

	for merged != nil {
		fmt.Println(merged.Value, " ")
		merged = merged.Next
	}

}

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func merageTwoLinkedList(l1 *LinkedList, l2 *LinkedList) *LinkedList {
	dummy := &LinkedList{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Value < l2.Value {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}

		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	}

	if l2 != nil {
		current.Next = l1
	}
	return dummy.Next
}
