package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	ch := make(chan int, 100)

	wg.Add(2)

	for i := 0; i < 100; i++ {
		ch <- i
	}

	go printEven(ch, &wg)
	go printOdd(ch, &wg)

	close(ch)
	wg.Wait()
}

func printEven(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case num, ok := <-ch:
			if !ok {
				fmt.Println("Even go routine is Done!!")
				return
			} else if num%2 == 0 {
				fmt.Printf("Even: %d\n", num)

			}

		}
	}
}

func printOdd(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case num, ok := <-ch:
			if !ok {
				fmt.Println("Odd go routine is Done !!")
				return
			} else if num%2 != 0 {
				fmt.Printf("Odd:%d\n", num)

			}

		}
	}
}
