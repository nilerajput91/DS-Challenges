package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {

	urls := []string{"https:dpsnashik.com", "https://targetrwe.com", "https://techracers1.greythr.com/"}
	var wg sync.WaitGroup

	resultChan := make(chan string, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go makeApiRequest(url, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for _, result :=range <-resultChan{
		fmt.Println(result)

	}

}

func makeApiRequest(url string, wg *sync.WaitGroup, resultChan chan<- string) {
	defer wg.Done()
	resp, err := http.Get(url)

	if err != nil {
		resultChan <- fmt.Sprintf("Error making request to %s:%v", url, err)
		return
	}

	defer resp.Body.Close()

	resultChan <- fmt.Sprintf("Response from %s:Status code:%d", resp, resp.StatusCode)

	
}
