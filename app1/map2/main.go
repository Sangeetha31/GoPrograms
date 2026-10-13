package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	courses := make(map[int]string)

	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Println("Enter no. of Elements:")
	fmt.Scan(&n)
	reader.ReadString('\n')

	for i := 1; i <= n; i++ {
		fmt.Printf("Enter Course-%d Name:", i)

		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		courses[i] = name
	}

	fmt.Println("The Map Values:")
	fmt.Println(courses)

}
