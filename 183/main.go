package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	for i := 1; i < 5; i++ {
		go producer(ch)
		fmt.Println("Counsuming:-->", <-ch)
		time.Sleep(500 * time.Millisecond)
	}

}

func producer(ch chan<- int) {
	for i := 1; i < 5; i++ {
		fmt.Println("producer", i)
		ch <- i
	}
	close(ch)
}
