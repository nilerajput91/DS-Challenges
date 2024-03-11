package main

import (
	"fmt"
	"slices"
)

func main() {

	in := []string{"ab"}

	fmt.Println(binarysearch(in))

}

func binarysearch(arr []string) (int, bool) {
	position, found := slices.BinarySearch(arr,"a")

	return position, found

}
