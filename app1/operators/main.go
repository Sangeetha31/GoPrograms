package main

import "fmt"

func main() {

	var i int = 10
	fmt.Println(i)

	i = 25
	fmt.Println(i)

	i += 10
	fmt.Println(i)

	i -= 10
	fmt.Println(i)

	i *= 10
	fmt.Println(i)

	i /= 10
	fmt.Println(i)

	var a int = 10
	var b int = 8

	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a == b)
	fmt.Println(a != b)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)

	fmt.Println(!true)
	fmt.Println(!false)

	var x int = 10
	fmt.Println(x)
	x++
	fmt.Println(x)
}
