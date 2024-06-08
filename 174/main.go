package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan *int, 4)
	array := []int{1, 2, 3, 4}

	var wg sync.WaitGroup

	wg.Add(len(array))

	go func() {

		for _, value := range array {
			v := value

			ch <- &v
		}

	}()

	go func() {
		for val := range ch {
			fmt.Println(*val)
			wg.Done()
		}

	}()

	wg.Wait()
}
