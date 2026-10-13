package main

import "fmt"

func main() {

	var n int

	fmt.Println("Enter a value")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Println("Hello")
	}

	i := 1
	for {
		fmt.Println(i)
		i++
		if i > n {
			break
		}
	}
}
