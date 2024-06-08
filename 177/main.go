package main

import (
	"fmt"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	input := []string{
		"I am so happy today!",
		"I hate this weather.",
		"Happy birthday!!",
	}

	resultChan := make(chan string)

	counter := 0

	for _, data := range input {
		wg.Add(1)

		go analyzeSentiment(data, resultChan, &counter)
	}

	go func() {

		wg.Wait()
		close(resultChan)

	}()

	for i := 0; i < len(input); i++ {
		fmt.Println(<-resultChan)
	}
	fmt.Printf("%d out of %d input strings had a positive sentiment.\n", counter, len(input))

}

func analyzeSentiment(data string, resultChan chan string, counter *int) {

	defer wg.Done()

	if strings.Contains(strings.ToLower(data), "happy") {
		mu.Lock()
		*counter++
		mu.Unlock()

		resultChan <- "positive"

	} else {
		resultChan <- "negative"
	}

}
