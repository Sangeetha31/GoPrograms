package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Goroutine Started:", id)

	for i := 1; i <= 5; i++ {
		fmt.Println("Goroutine", id, ":", i)
		time.Sleep(time.Second)
	}

	fmt.Println("Goroutine Finished:", id)
}
func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go worker(1, &wg)
	go worker(2, &wg)

	wg.Wait()

	fmt.Println("Main Execution Complete")
}
