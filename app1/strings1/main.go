package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string
	s1 = "india"
	fmt.Println(s1)

	var s2 string = "india"
	fmt.Println(s2)

	s3 := "india"
	fmt.Println(s3)

	firstName := "Ram"
	lastName := "Murthy"
	fullName := firstName + " " + lastName
	fmt.Println(fullName)

	fmt.Println(s1 == s2)
	fmt.Println(s1 != s2)
	fmt.Println(firstName <= lastName)

	fmt.Println(strings.ToUpper(s1))
	fmt.Println(strings.ToLower(fullName))
	fmt.Println(strings.TrimSpace("        INDIA             "))

	str1 := "i am from INDIA"
	fmt.Println(strings.Replace(str1, "m", "$", 3))
	fmt.Println(strings.Replace(str1, "m", "$", 1))
	fmt.Println(strings.Replace(str1, "from", "in", 1))

	fmt.Println(strings.ReplaceAll(str1, "m", "$"))

	arr := []string{"Ram", "Harry", "Priya", "Krishna"}
	fmt.Println(strings.Join(arr, "|"))

	fmt.Println(strings.Split("Ram|Harry|Priya|Krishna", "|"))

	fmt.Println(strings.HasPrefix(str1, "i am"))
	fmt.Println(strings.HasSuffix(str1, "INDIA"))

	str2 := "i am from INDIA"
	str3 := "i am from india"

	fmt.Println(str3 == str2)
	fmt.Println(strings.EqualFold(str2, str3))
	fmt.Println(len(str3))
	fmt.Println(strings.Compare(str2, str3))

}
