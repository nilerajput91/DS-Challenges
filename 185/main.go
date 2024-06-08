package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	data := &Data{Value: "Intial"}

	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go modfiyData(data, &wg)
	}
	wg.Done()

	fmt.Println("All goroutine Done..")

}

type Data struct {
	sync.Mutex
	Value string
}

func modfiyData(data *Data, wg *sync.WaitGroup) {
	defer wg.Done()

	data.Lock()
	data.Value = data.Value
	data.Unlock()

	time.Sleep(200 * time.Second)

	data.Lock()
	data.Value = data.Value + "modified"
	data.Unlock()
	fmt.Println("modified data by goroutine", data.Value)

}
