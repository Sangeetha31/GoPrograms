package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	var wg sync.WaitGroup
	var mut sync.Mutex

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mut.Lock()
			count++
			mut.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
