package main

import "fmt"

type ListNode struct {
	value int
	Next  *ListNode
}
func rearrangeLinkedList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    // Separate the list into odd and even nodes
    oddHead := head
    evenHead := head.Next
    oddTail := head
    evenTail := head.Next

    // Traverse the list to separate odd and even nodes
    current := evenHead.Next
    isOdd := true
    for current != nil {
        if isOdd {
            oddTail.Next = current
            oddTail = current
        } else {
            evenTail.Next = current
            evenTail = current
        }
        isOdd = !isOdd
        current = current.Next
    }

    // Connect the two separated lists
    oddTail.Next = evenHead
    evenTail.Next = nil

    return oddHead
}


func PrintlinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print( current.value)
		current = current.Next
	}
}

func main() {
	head1 := &ListNode{value: 1, Next: &ListNode{value: 2, Next: &ListNode{value: 3, Next: &ListNode{value: 4}}}}
	result := rearrangeLinkedList(head1)
	PrintlinkedList(result)
}
