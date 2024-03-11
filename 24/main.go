package main

import "fmt"

func main() {

	myLinkedList := LinkedList{}
	myLinkedList.Appedn(1)
	myLinkedList.Appedn(2)
	myLinkedList.Appedn(3)
	fmt.Println("Linked List: ")
	myLinkedList.Print()

}

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Appedn(data int) {
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

func (list *LinkedList) Print() {
	current := list.Head
	for current != nil {
		fmt.Printf("%d ->", current.Data)
		current = current.Next

	}
	fmt.Println("nil")
}
