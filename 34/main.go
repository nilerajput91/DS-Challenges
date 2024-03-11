package main

import "fmt"

func main() {
	myStack := Stack{}
	pushItem := myStack.Push(10)
	fmt.Println(pushItem)

	IsEmpty := myStack.IsEmpty()
	fmt.Println(IsEmpty)

}

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) []int {
	s.items = append(s.items, item)
	return s.items
}

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, fmt.Errorf("stack is empty")

	}

	item := s.items[len(s.items)-1]

	fmt.Println(item)

	s.items = s.items[:len(s.items)-1]

	return item, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.items) == 0 {
		return 0, fmt.Errorf("stack is empty no peek to show")
	}

	return s.items[len(s.items)-1], nil

}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0

}
