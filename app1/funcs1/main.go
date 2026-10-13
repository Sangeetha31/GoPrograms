package main

import "fmt"

func getSum(nums ...int){
	total:=0
}

for _,num:=range nums{
	tota+=num
}
fmt.Println("The sum is:",total)
}

func main(){
	getSum(10,20,36,49)
	getSum(15,88,99,77,55,66)

	fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Three")
	fmt.Println("Four")
}