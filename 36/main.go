package main

import "fmt"

func main() {

	deq := &Dequeue{}

	

	deq.PushFront(1)
	deq.PushRear(2)
	empty := deq.IsEmpty()
	fmt.Println(empty)

	ele,_:=deq.PopFront()
	fmt.Println("Pop Front Element:",ele)

	len:=deq.Size()
	fmt.Println("size of dequeue:",len)

}

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

type Dequeue struct {
	Front *Node
	Rear  *Node
}

func (d *Dequeue) PushFront(item int) {
	newNode := &Node{
		Data: item,
		Next: nil,
		Prev: nil,
	}

	if d.Front == nil {
		d.Front = newNode
		d.Rear = newNode
	} else {
		newNode.Next = d.Front
		d.Front.Prev = newNode
		d.Front = newNode
	}
}

func (d *Dequeue) PushRear(item int) {
	newNode := &Node{
		Data: item,
		Next: nil,
		Prev: nil,
	}

	if d.Rear == nil {
		d.Rear = newNode
		d.Front = newNode
	} else {
		newNode.Prev = d.Rear
		d.Rear.Next = newNode
		d.Rear = newNode
	}

}

func (d *Dequeue) PopFront() (int, error) {
	if d.Front == nil {
		return 0, fmt.Errorf("dequeue is empty")
	}

	item := d.Front.Data

	if d.Front == d.Rear {
		d.Front = nil
		d.Rear = nil
	} else {
		d.Front = d.Front.Next
		d.Front.Prev = nil
	}

	return item, nil
}

func (d *Dequeue) PopRear() (int, error) {
	if d.Rear == nil {
		return 0, fmt.Errorf("dequeue is empty")
	}

	item := d.Rear.Data

	if d.Front == d.Rear {
		d.Front = nil
		d.Rear = nil

	} else {
		d.Rear = d.Rear.Prev
		d.Rear.Next = nil
	}

	return item, nil
}

func (d *Dequeue) IsEmpty() bool {
	return d.Front == nil
}

func (d *Dequeue) Size() int {
	size := 0
	current := d.Front

	for current != nil {
		size++
		current = current.Next
	}

	return size

}
