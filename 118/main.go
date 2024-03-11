/* Given a singly linked list of size N containing only English Alphabets. Your task is to complete the function arrangeC&V(), that arranges the consonants and vowel nodes of the list it in such a way that all the vowels nodes come before the consonants while maintaining the order of their arrival.

Input:
The function takes a single argument as input, the reference pointer to the head of the linked list. There will be T test cases and for each test case the function will be called separately.

Output:
For each test case output a single line containing space separated elements of the list.

User Task:
The task is to complete the function arrange() which should arrange the vowels and consonants as required.

Constraints:
1 <= T <= 100
1 <= N <= 100

Example:
Input:
2
6
a e g h i m
3
q r t

Output:
a e i g h m
q r t

Explanation:
Testcase 1: Vowels like a, e and i are in the front, and consonants like g, h and m are at the end of the list. */

package main

import "fmt"

func main() {
	head := &Node{Data: 'a'}
	head.Next = &Node{Data: 'e'}
	head.Next.Next = &Node{Data: 'g'}
	head.Next.Next.Next = &Node{Data: 'h'}
	head.Next.Next.Next.Next = &Node{Data: 'i'}
	head.Next.Next.Next.Next.Next = &Node{Data: 'm'}

	fmt.Println("-----")
	head = arrangecv(head)
	printList(head)

}

type Node struct {
	Data int
	Next *Node
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func arrangecv(head *Node) *Node {
	if head == nil {
		return nil
	}

	// create separate lists for vowels and consonats
	var vowelHead, vowelTail, consonantHead, consonantTail *Node

	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = nil

		if isVowel(byte(head.Data)) {
			if vowelHead == nil {
				vowelHead = curr
				vowelTail = curr

			} else {
				vowelTail.Next = curr
				vowelTail = curr
			}
		} else {
			if consonantHead == nil {
				consonantHead = curr
				consonantTail = curr
			} else {
				consonantTail.Next = curr
				consonantTail = curr
			}
		}
		curr = next

	}

	if vowelHead == nil {
		return consonantHead
	}
	vowelTail.Next = consonantHead
	return vowelHead

}

func printList(head *Node) {
	for head != nil {
		fmt.Printf("%s \n", string(head.Data))
		head = head.Next
	}
}
