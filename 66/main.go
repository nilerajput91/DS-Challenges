package main

import (
	"fmt"
	"sync"
)

func main() {

	numbers := []int{1, 3, 4, 5, 6, 7, 8}

	resultChan := make(chan int, len(numbers))

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(2)

		go squareCal(num, resultChan, &wg)
		go squareCal(num, resultChan, &wg)

	}

	go func() {
		wg.Wait()

		close(resultChan)
	}()

	for result := range resultChan {
		fmt.Println(result)
	}

}

func squareCal(num int, resultChan chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	square := num * num

	resultChan <- square
}
