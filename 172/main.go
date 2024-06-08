//Semaphore Pattern:

/*
	The semaphore pattern uses a channel to limit the number of goroutines that can run concurrently.

You can create a buffered channel with a capacity equal to the desired limit.
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	numWorkers := 3
	sem := make(chan struct{}, numWorkers)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, sem, &wg)
	}
	wg.Wait()
}

func worker(id int, sem chan struct{}, wg *sync.WaitGroup) {

	defer wg.Done()

	sem <- struct{}{}

	defer func() {
		<-sem
	}()

	fmt.Printf("Worker %d is working\n", id)
}
