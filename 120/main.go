/* Arya has a string, s, of uppercase English letters. She writes down the string s on a paper k times. She wants to calculate the occurence of a specific letter  in the first n characters of the final string.

Example

Input :
s = "ABA"
k = 3
n = 7
c = 'B'
Output :
2
Explaination :
Final string - ABAABAABA
Example

Input :
s = "MMM"
k = 2
n = 4
c = 'M'
Output :
4
Explaination :
Final string - MMMMMM
â€‹Your Task:
You don't need to read input or print anything. Your task is to complete the function fun() which takes the string s , integer k ,integer n and the character c as input parameter and returns the occurence of the character c  in the first n characters of the final string.

Expected Time Complexity : O(|s|)
Expected Auxiliary Space : O(1) */

package main

import "fmt"

func main() {
	s := "ABA"
	k := 3
	n := 7
	c := 'B'

	fmt.Printf("occurances of char :%d\n", fun(s, k, n, byte(c)))

}

func fun(s string, k, n int, c byte) int {

	// len of the original string
	strLen := len(s)

	// cal the number of time s is repeated

	repeat := n / strLen

	// cal the remaing character the full repetation

	remain := n % strLen

	// count the full occurances of the char c
	count := 0

	for i := 0; i < strLen; i++ {
		if s[i] == c {
			count += repeat
		}

	}

	for j := 0; j < remain; j++ {
		if s[j] == c {
			count++
		}
	}
	return count

}
