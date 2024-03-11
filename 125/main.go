/*
	Given an unsorted linked list of N LinkedLists. The task is to remove duplicate elements from this unsorted Linked List. When a value appears in multiple LinkedLists, the LinkedList which appeared first should be kept, all others duplicates are to be removed.

Example 1:

Input:
N = 4
value[] = {5,2,2,4}
Output: 5 2 4
Explanation:Given linked list elements are
5->2->2->4, in which 2 is repeated only.
So, we will delete the extra repeated
elements 2 from the linked list and the
resultant linked list will contain 5->2->4
Example 2:

Input:
N = 5
value[] = {2,2,2,2,2}
Output: 2
Explanation:Given linked list elements are
2->2->2->2->2, in which 2 is repeated. So,
we will delete the extra repeated elements
2 from the linked list and the resultant
linked list will contain only 2.
Your Task:
You have to complete the method removeDuplicates() which takes 1 argument: the head of the linked list.  Your function should return a pointer to a linked list with no duplicate element.
Expected Time Complexity: O(N)
Expected Auxiliary Space: O(N)
Constraints:
1 <= size of linked lists <= 106
0 <= numbers in list <= 104
*/
package main

import "fmt"

func main() {

	head := &LinkedList{
		Data: 2,
	}
	head.Next = &LinkedList{
		Data: 1,
	}
	head.Next.Next = &LinkedList{
		Data: 2,
	}

	head.Next.Next.Next = &LinkedList{
		Data: 2,
	}

	head = removeDuplicate(head)

	for head != nil {
		fmt.Printf("%d \n", head.Data)
		head = head.Next
	}

}

type LinkedList struct {
	Data int
	Next *LinkedList
}

func removeDuplicate(head *LinkedList) *LinkedList {
	// Map to store the last occurrence index of each element
	lastOccurrence := make(map[int]int)

	// Iterate through the linked list to find the last occurrence of each element
	curr := head
	index := 0
	for curr != nil {
		lastOccurrence[curr.Data] = index
		curr = curr.Next
		index++
	}

	// Reconstruct the linked list using the map
	dummy := &LinkedList{Data: -1}
	prev := dummy
	for i := 0; i < index; i++ {
		if lastIdx, ok := lastOccurrence[i]; ok {
			curr := head
			for j := 0; j < lastIdx; j++ {
				curr = curr.Next
			}
			prev.Next = &LinkedList{Data: curr.Data}
			prev = prev.Next
		}
	}

	return dummy.Next

}
