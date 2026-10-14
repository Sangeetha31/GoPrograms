package main

import "fmt"

func main() {

	ch := make(chan int, 5)

	ch <- 125
	ch <- 165
	ch <- 228
	ch <- 985
	ch <- 658

	<-ch
	fmt.Println(<-ch)

	arr := []int{
		<-ch, <-ch, <-ch, <-ch, <-ch,
	}
	fmt.Println(arr[1])
}
