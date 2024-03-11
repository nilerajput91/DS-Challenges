package main

import "fmt"

func main() {
	str := "abb"
	result := isPalidrome(str)
	fmt.Println(result)

}

//Recursive

func isPalidrome(str string) bool {
	//runeStr:=[]rune(str)
	start := 0

	end := len(str)-1

	if start >= end {
		return true
	}

	return (str[start] == str[end]) && isPalidrome(str)

}
