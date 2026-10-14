package main

import (
	"fmt"
	"reflect"
)

type Product struct {
	id       int
	name     string
	price    float32
	category string
}

func main() {
	i := 15

	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.ValueOf(i))

	str := "india"
	val := reflect.ValueOf(str)
	fmt.Println(val.Type())
	fmt.Println(val.Kind())
	fmt.Println(val.String())

	pr := Product{id: 121, name: "Dell 15 3525", price: 48900, category: "Laptops"}

	val1 := reflect.ValueOf(pr)
	typ := reflect.TypeOf(pr)

	for i := 0; i < val1.NumField(); i++ {
		fmt.Println(typ.Field(i).Name, "-", val1.Field(i))
	}

}
