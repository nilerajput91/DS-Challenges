package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	inputChan := make(chan string, 10)
	resultChan := make(chan string, 10)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go worker(inputChan, resultChan, i)

	}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		inputChan <- line
	}
	close(inputChan)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	numPositive := 0
	numNegative := 0

	for result := range resultChan {

		switch result {
		case "positive":
			numPositive++

		case "negative":
			numNegative++
		}
	}

	fmt.Printf("Positive:%d\n", numPositive)
	fmt.Printf("Negative:%d\n", numNegative)

}

func analyzeSentiment(data string, resultChan chan<- string) {
	if strings.Contains(strings.ToLower(data), "happy") {

		resultChan <- "positive"
	} else {
		resultChan <- "negative"
	}
}

func worker(inputChan <-chan string, resultChan chan<- string, k int) {
	defer wg.Done()

	for data := range inputChan {
		go analyzeSentiment(data, resultChan)

		mu.Lock()
		fmt.Printf("worker %d processed line:%s\n", k, data)
		mu.Unlock()
	}
}
