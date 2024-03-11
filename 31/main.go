package main

import "fmt"

func main() {
	mylist := &LinkedList{}
	mylist.Append(1)
	mylist.Append(2)
	mylist.Append(3)

	fmt.Println("original linked list ")
	mylist.PrintList()

	mylist.reverse()

	fmt.Println("reverse linked list ")
	mylist.PrintList()
}

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Append(data int) {
	newNode := &Node{
		Data: data,
		Next: nil,
	}

	if list.Head == nil {
		list.Head = newNode
		return

	}

	current := list.Head
	if current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (list *LinkedList) reverse() {

	var prev, current, next *Node
	current = list.Head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	list.Head = prev

}

func (list *LinkedList) PrintList() {
	current := list.Head
	for current != nil {
		fmt.Println(current.Data, " ")
		current = current.Next
	}
	fmt.Println()
}
