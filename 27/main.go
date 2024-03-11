package main

import "fmt"

func main() {

	str1 := "ABCDEF"
	str2 := "AED"

	isSubseq := isSubseqence(str1,str2,6,3)
	fmt.Println(isSubseq)

}

func isSubseqence(str1, str2 string, lenstr1, lenstr2 int) bool {

	j := 0

	for i := 0; i < lenstr1 && j < lenstr2; i++ {
		if str1[i] == str2[j] {
			j++
		}
	}

	return j == lenstr2

}
