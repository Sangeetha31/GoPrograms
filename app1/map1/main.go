package main

import "fmt"

func main() {

	courses := map[int]string{
		1201: "Java",
		1202: "React",
		1209: "SQL",
		1325: "GO",
	}

	for id := range courses {
		fmt.Println(id)
	}
	for _, name := range courses {
		fmt.Println(name)
	}

	for id, name := range courses {
		fmt.Println(id, name)
	}
}
