package main

import (
	"fmt"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	input := []string{
		"I am so happy today!",
		"I hate this weather.",
		"Happy birthday!!",
	}

	resultchan := make(chan string)

	for _, data := range input {
		wg.Add(1)
		go analyzeSentiment(data, resultchan)
	}

	go func() {
		wg.Wait()
		close(resultchan)

	}()

	for i := 0; i < len(input); i++ {
		fmt.Println(<-resultchan)
	}

}
func analyzeSentiment(data string, resultChan chan string) {
	defer wg.Done()

	if strings.Contains(strings.ToLower(data), "happy") {

		resultChan <- "positive"
	} else {
		resultChan <- "negative"
	}
}
