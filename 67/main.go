package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	oddEvenChanResult := make(chan int)

	for i := 0; i < 100; i++ {
		wg.Add(2)

		go printEvenOdd(i, oddEvenChanResult, &wg)
		go printEvenOdd(i, oddEvenChanResult, &wg)

	}

	go func() {
		wg.Wait()
		close(oddEvenChanResult)
	}()

	for num := range oddEvenChanResult {
		fmt.Printf("2 go routine result %d \n", num)
	}
}

func printEvenOdd(num int, evenOddChan chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	if num%2 == 0 {
		evenOddChan <- num
	} else {
		evenOddChan <- num

	}

}
