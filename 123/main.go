/* Given a linked list and a number k. You have to reverse first part of linked list with k nodes and the second part with n-k nodes.

Example 1:

Input: 1 -> 2 -> 3 -> 4 -> 5
k = 2
Output: 2 -> 1 -> 5 -> 4 -> 3
Explanation: As k = 2 , so the first part 2
nodes: 1 -> 2 and the second part with 3 nodes:
3 -> 4 -> 5. Now after reversing the first part:
2 -> 1 and the second part: 5 -> 4 -> 3.
So the output is: 2 -> 1 -> 5 -> 4 -> 3
Example 2:

Input: 1 -> 2 -> 4 -> 3
k = 3
Output: 4 -> 2 -> 1 -> 3
Explanation: As k = 3 , so the first part
3 nodes: 4 -> 2 -> 1 and the second part
with 1 nodes: 3. Now after reversing the
first part: 1 -> 2 -> 4 and the
second part: 3. So the output is: 1 -> 2 -> 4 -> 3


Your Task:
You don't need to read input or print anything. Your task is to complete the function reverse() which takes head node of the linked list and a integer k as input parameters and returns head node of the linked list after reversing both parts.


Constraints:
1 <= N <= 105
1 <= k < N

Expected Time Complexity: O(N)
Expected Space Complexity: O(1) */

package main

import "fmt"

func main() {
	head := &Node{
		Data: 1,
	}
	head.Next = &Node{
		Data: 2,
	}
	head.Next.Next = &Node{
		Data: 3,
	}
	head.Next.Next.Next = &Node{
		Data: 4,
	}
	head.Next.Next.Next.Next = &Node{
		Data: 5,
	}

	k := 3

	head = reverse(head, k)

	if head != nil {
		for head != nil {
			fmt.Printf("%d\n", head.Data)
			head = head.Next
		}
	}

}

type Node struct {
	Data int
	Next *Node
}

func reverse(head *Node, k int) *Node {
	if !hasNode(head, k) {
		return head
	}

	var prev *Node
	count := 0
	curr := head

	for curr != nil && count < k {
		curr = curr.Next
		count++
	}

	if count == k {
		prev = reverseList(head, k)
		head.Next = reverse(curr, k)

	}
	return prev

}

func reverseList(head *Node, k int) *Node {
	var prev, next *Node
	count := 0

	curr := head

	for curr != nil && count < k {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		count++
	}
	return prev
}

func hasNode(head *Node, k int) bool {
	count := 0
	curr := head
	for curr != nil {
		count++
		curr = curr.Next
	}
	return count >= k

}
