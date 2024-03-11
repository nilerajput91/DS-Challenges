/* Given an array of n names arr of candidates in an election, where each name is a string of lowercase characters. A candidate name in the array represents a vote casted to the candidate. Print the name of the candidate that received the maximum count of votes. If there is a draw between two candidates, then print lexicographically smaller name.

Example 1:

Input:
n = 13
arr[] = {john,johnny,jackie,johnny,john
jackie,jamie,jamie,john,johnny,jamie,
johnny,john}
Output: john 4
Explanation: john has 4 votes casted for
him, but so does johny. john is
lexicographically smaller, so we print
john and the votes he received.
Example 2:

Input:
n = 3
arr[] = {andy,blake,clark}
Output: andy 1
Explanation: All the candidates get 1
votes each. We print andy as it is
lexicographically smaller.
Your Task:
You only need to complete the function winner() that takes an array of strings arr, and length of arr n as parameters and returns an array of string of length 2. First element of the array should be the name of the candidate and second element should be the number of votes that candidate got in string format.

Expected Time Complexity: O(n)
Expected Auxiliary Space: O(n)

Constraints:
1 <= n <= 105
1 <= |arri| <= 105 */

package main

import "fmt"

func main() {
	arr := []string{"andy", "blake", "clark"}
	n := len(arr)

	resultWinner := winner(arr, n)
	fmt.Printf("winner: %v,count: %v \n", resultWinner[0], resultWinner[1])

}

func winner(str []string, num int) []string {
	votes := make(map[string]int)

	for _, name := range str {
		votes[name]++
	}

	maxVotes := 0
	winner := ""

	for name, count := range votes {

		if count > maxVotes || (count == maxVotes && name < winner) {
			maxVotes = count
			winner = name
		}

	}

	return []string{winner, fmt.Sprintf("%d", maxVotes)}

}
