package main

import "fmt"

func main() {


	names := []string{ "Geek", "Geeks", "Geeksfor",
  "GeeksforGeek", "GeeksforGeeks" }

  fmt.Println( longestName(names) )

}

func longestName(names []string) string {

	if len(names) == 0 {

		return ""
	}

	longestName := names[0]
	maxLength := len(longestName)

	for _, name := range names {

		currentlength := len(name)
		if currentlength > maxLength {
			longestName = name
			maxLength = currentlength

		}
	}

	return longestName

}
