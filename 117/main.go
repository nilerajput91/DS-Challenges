/* Given a singly linked list L0 -> L1 -> … -> Ln-1 -> Ln. Rearrange the nodes in the list so that the new formed list is: L0 -> Ln -> L1 -> Ln-1 -> L2 -> Ln-2.

Input:
You have to complete the method which takes 1 argument: the head of the  linked list. You should not read any input from stdin/console. There are multiple test cases. For each test case, this method will be called individually.

Output:
Your function should return a pointer to the rearranged list so obtained.

User Task:
The task is to complete the function inPlace() which should rearrange the given linked list as required.

Constraints:
1 <=T<= 50
1 <= size of linked lists <= 100

Example:
Input:
2
4
1 2 3 4
5
1 2 3 4 5

Output:
1 4 2 3
1 5 2 4 3

Explanation:
Testcase 1: After rearranging the linked list as required, we have 1, 4, 2 and 3 as the elements of the linked list.


*/

package main

import "fmt"

func main() {
	  // Example usage
	  head1 := &Node{Data: 1}
	  head1.Next = &Node{Data: 2}
	  head1.Next.Next = &Node{Data: 3}
	  head1.Next.Next.Next = &Node{Data: 4}


	  fmt.Println("orignial linkedlist:")

	  printList(head1)

	  head1=inPlace(head1)

	  fmt.Println("rearrange linked list:")
	  fmt.Println()

	  printList(head1)


}

type Node struct {
	Data int
	Next *Node
}

func reverseList(head *Node) *Node {
	var prev, next *Node
	curr := head

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func findMiddle(head *Node) *Node {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

	}
	return slow

}

func mergeList(first, second *Node) *Node {
	dummy := &Node{}
	curr := dummy

	for first != nil && second != nil {
		curr.Next = first
		first = first.Next

		curr = curr.Next

		curr.Next = second
		second = second.Next
		curr = curr.Next

	}

	if first != nil {
		curr.Next = first
	}
	return dummy.Next
}
func inPlace(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	middle := findMiddle(head)
	secondHalf := reverseList(middle.Next)
	middle.Next = nil

	return mergeList(head, secondHalf)
}

func printList(head *Node) {
	for head != nil {
		fmt.Printf("%d ", head.Data)
		head = head.Next
	}
}
