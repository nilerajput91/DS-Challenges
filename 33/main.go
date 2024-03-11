package main

import (
	"fmt"
)

func main() {

	myCircularLinkedList := CircularLinkedList{}
	myCircularLinkedList.Insert(1)
	myCircularLinkedList.Insert(2)
	myCircularLinkedList.InsertAtFront(0)
	myCircularLinkedList.PrintList()

}

type Node struct {
	Data int
	Next *Node
}

type CircularLinkedList struct {
	Head *Node
}

func (cl *CircularLinkedList) Insert(data int) {
	newNode := &Node{
		Data: data,
		Next: nil,
	}

	if cl.Head == nil {
		cl.Head = newNode
		newNode.Next = newNode
		return
	}

	current := cl.Head

	for current.Next != cl.Head {
		current = current.Next
	}
	current.Next = newNode
	newNode.Next = cl.Head
}

func (list CircularLinkedList) InsertAtFront(data int) {
	newNode := &Node{Data: data,
		Next: nil}

	if list.Head == nil {
		list.Head = newNode
		newNode.Next = newNode
		return
	}

	lastNode := list.Head

	for lastNode.Next != list.Head {
		lastNode = lastNode.Next

	}

	lastNode.Next = newNode
	newNode.Next = list.Head

	list.Head = newNode.Next
}

func (cl *CircularLinkedList) PrintList() {
	if cl.Head == nil {
		fmt.Println("Circular linked List is empty ")
		return
	}

	current := cl.Head

	for {

		fmt.Println(current.Data, " ")
		current = current.Next

		if current == cl.Head {

			break
		}

	}
	fmt.Println()
}
