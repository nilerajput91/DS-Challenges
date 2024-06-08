/* Worker Pool:

A worker pool maintains a fixed number of goroutines that are continuously available
to process tasks from a queue. This approach is useful when you have a
stream of tasks that need to be processed concurrently but limited by the number of available workers. */

package main

import (
	"fmt"
	"sync"
)

func main() {
	numWorkers := 3
	numJobs := 10
	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()

}

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d is processing job %dd\n", id, job)

	}
}
