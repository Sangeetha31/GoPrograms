package main

import (
	"fmt"
	"time"
)

func print(id int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("GoRoutine %d: %d\n", id, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go print(1)
	go print(2)

	time.Sleep(6 * time.Second)
	fmt.Println("Execution Complete")
}
