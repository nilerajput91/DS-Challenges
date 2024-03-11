/* Given a sorted linked list, delete all nodes that have duplicate numbers (all occurrences), leaving only numbers that appear once in the original list.

Example 1:

Input:
N = 8
Linked List = 23->28->28->35->49->49->53->53
Output:
23 35
Explanation:
The duplicate numbers are 28, 49 and 53 which
are removed from the list.
Example 2:

Input:
N = 6
Linked List = 11->11->11->11->75->75
Output:
Empty list
Explanation:
All the nodes in the linked list have
duplicates. Hence the resultant list
would be empty.
Your task:
You don't have to read input or print anything. Your task is to complete the function removeAllDuplicates() which takes the head of the linked list, removes all the occurences of duplicates in the linked list and returns the head of the modified linked list.

Expected Time Complexity: O(N)
Expected Auxiliary Space: O(1)
*/

package main

import "fmt"

func main() {

	head := &Node{
		Data: 23,
	}
	head.Next = &Node{
		Data: 28,
	}

	head.Next.Next = &Node{
		Data: 28,
	}

	head.Next.Next.Next = &Node{
		Data: 35,
	}

	head.Next.Next.Next.Next = &Node{
		Data: 49,
	}
	head.Next.Next.Next.Next.Next = &Node{
		Data: 53,
	}
	head.Next.Next.Next.Next.Next.Next = &Node{
		Data: 53,
	}

	head = removeAllDuplicates(head)

	for head != nil {
		fmt.Printf("%d \n", head.Data)
		head = head.Next
	}
}

type Node struct {
	Data int
	Next *Node
}

func removeAllDuplicates(head *Node) *Node {

	dummy := &Node{
		Data: -1,
		Next: head,
	}

	prev := dummy
	curr := head

	for curr != nil {
		for curr.Next != nil && curr.Data == curr.Next.Data {
			curr = curr.Next
		}

		if prev.Next != curr {
			prev.Next = curr.Next
		} else {
			prev = prev.Next
		}

		curr = curr.Next

	}
	return dummy.Next
}
