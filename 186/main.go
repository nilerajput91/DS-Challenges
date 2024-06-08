package main

import (
	"fmt"
	"sync"
)

func main() {
	chLower := make(chan bool, 1)
	chUpper := make(chan bool, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	go printLower(chLower, chUpper, &wg)
	go printUpper(chLower, chUpper, &wg)
	chLower <- true
	wg.Wait()

	close(chLower)
	close(chUpper)

}

func printLower(chLower, chUpper chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'a'; i <= 'z'; i++ {
		<-chLower
		fmt.Printf("%c", i)
		chUpper <- true

	}
}
func printUpper(chLower, chUpper chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'A'; i <= 'Z'; i++ {
		<-chUpper
		fmt.Printf("%c", i)
		chLower <- true
	}
}
