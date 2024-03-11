package main

import "fmt"

func main() {

	mylist:=LinkedList{}

	mylist.Appnd(1)
	mylist.Appnd(1)
	mylist.Appnd(2)
	mylist.Appnd(3)
	mylist.Appnd(3)
	mylist.Appnd(4)
	mylist.PrintList()

	mylist.removeDuplicateList()
	mylist.PrintList()
}

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Appnd(data int) {

	newNode := &Node{
		Data: data,
		Next: nil,
	}

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

func (list *LinkedList) removeDuplicateList() {
	current := list.Head

	for current != nil && current.Next != nil {
		if current.Data == current.Next.Data {
			current.Next = current.Next.Next
		} else {

			current = current.Next
		}
	}
}

func (list *LinkedList) PrintList() {
	current := list.Head

	for current.Next != nil {
		fmt.Println("List:", current.Data)
		current = current.Next
	}

	fmt.Println()
}
