package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hey User! This Message is Through Channel"
	}()

	msg := <-ch

	fmt.Println(msg)
}
