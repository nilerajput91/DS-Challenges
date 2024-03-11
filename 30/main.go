package main

import (
	"fmt"
)

func main() {
	myList := &LinkedList{}
	myList.Append(1)
	myList.Append(1)

	myList.PrintList()
	//myList.DeleteFirst()
	myList.DeleteLast()
	myList.PrintList()

}

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Append(data int) {

	newNode := &Node{Data: data, Next: nil}

	if list.Head == nil {
		list.Head = newNode
		return
	}

	current := list.Head

	for current.Next != nil {
		current = current.Next

	}

	current.Next = newNode

}

func (list *LinkedList) DeleteFirst() {
	if list.Head == nil {
		return
	}

	list.Head = list.Head.Next
}

func (list *LinkedList) PrintList() {
	current := list.Head
	for current != nil {
		fmt.Printf(" %d ->", current.Data)
		current = current.Next
	}

	fmt.Println()

}

func (list *LinkedList) DeleteLast() {
	if list.Head == nil {
		return
	}

	if list.Head.Next == nil {
		list.Head = nil
		return
	}

	current:=list.Head
	for current.Next.Next!=nil{
		current=current.Next
	}

	current.Next=nil
}
