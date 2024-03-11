package main

import "fmt"

func main() {

	myLinkedList := LinkedList{}
	myLinkedList.Append(1)
	myLinkedList.Append(2)
	
	myLinkedList.Print()

	bool:=myLinkedList.Search(3)
	fmt.Println(bool)

	myLinkedList.InsertAt(4,1)
	myLinkedList.Print()



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

	

	for current.Next != nil {
		current = current.Next
	}

	

	current.Next = newNode
}

func (list *LinkedList) Prepend(data int) {
	newNode := &Node{
		Data: data,
		Next: list.Head,
	}
	list.Head = newNode

}

func (list *LinkedList) InsertAt(data, pos int) {
	newNode := &Node{Data: data, Next: nil}
	if pos <= 0 {
		list.Prepend(data)
		return
	}

	current := list.Head

	for i := 0; i < pos-1 && current != nil; i++ {
		current = current.Next
	}

	if current == nil {
		list.Append(data)
		return
	}

	newNode.Next = current.Next
	current.Next = newNode
}

func (list *LinkedList) Delete(data int) {

	if list.Head == nil {
		return
	}

	if list.Head.Data == data {
		list.Head = list.Head.Next
		return
	}

	current := list.Head
	for current.Next != nil && current.Next.Data == data {
		current = current.Next
	}

	if current.Next == nil {
		return
	}

	current.Next = current.Next.Next
}

func (list *LinkedList) Search(data int) bool {
	current := list.Head

	for current != nil {
		if current.Data == data {
			return true
		}
		current = current.Next
	}
	return false
}

func (list *LinkedList) Print() {
	current := list.Head

	for current != nil {
		fmt.Printf(" %d ->", current.Data)
		current = current.Next
	}

	fmt.Println("nil")
}
