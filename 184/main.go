package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go longRunningtask(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Operation is cancelled:", ctx.Err())
	}

}
func longRunningtask(ctx context.Context) {

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Task completed successfully....")
	case <-ctx.Done():
		fmt.Println("Task cancelled:", ctx.Err())
	}
}
