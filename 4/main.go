// replace word with AB with c
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {
	inputstring := "helloABworld"
	outputString := replaceWord(inputstring)

	fmt.Println(outputString)

}

func replaceWord(str string) string {
	return strings.Replace(str, "AB", "C", -1)
}
