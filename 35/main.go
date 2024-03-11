package main

import "fmt"

func main() {

	myqueue := &Queue{}

	myqueue.Enqueue(1)
	myqueue.Enqueue(1)
	myqueue.Enqueue(1)

	fmt.Println(myqueue)

	result, _ := myqueue.Dequeue()
	fmt.Println(result)
	isEmpty := myqueue.IsEmpty()
	fmt.Println("Is Empty: ", isEmpty)
	totalElement := myqueue.Size()
	fmt.Println("total Queue:", totalElement)

	frontEle, _ := myqueue.Front()
	fmt.Println("Current Element::", frontEle)

}

type Queue struct {
	Items []interface{}
}

func (q *Queue) Enqueue(data int) {
	q.Items = append(q.Items, data)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if len(q.Items) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}

	items := q.Items[0]
	q.Items = q.Items[1:]

	return items, nil
}

func (q *Queue) IsEmpty() interface{} {
	return len(q.Items) == 0
}

func (q *Queue) Size() int {
	return len(q.Items)
}

func (q *Queue) Front() (interface{}, error) {
	if len(q.Items) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}

	return q.Items[0], nil

}
