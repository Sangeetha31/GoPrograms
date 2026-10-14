package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "UAT/129.52.654.25/admin/#12181hbvyub"
	}()

	go func() {
		ch2 <- "PRODUCTION/129.52.600.22/admin/#12181hbvyub"
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	}
}
