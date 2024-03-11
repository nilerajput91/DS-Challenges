package main

func main() {

}

func leftMostRepeatingChar(s string) int {
	charIndex := make(map[rune]int)

	leftMostIndexChar := -1

	for index, char := range s {

		if in, found := charIndex[char]; found {
			if leftMostIndexChar == -1 || in < leftMostIndexChar {
				leftMostIndexChar = in
			}
		} else {
			charIndex[char] = index
		}
	}

	return leftMostIndexChar
}
