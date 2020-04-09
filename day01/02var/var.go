package main

import "fmt"

var (
	name string
	age  int
	isOk bool
)

func main() {
	s1 := "ck"
	name = "曹坤"
	age = 23
	isOk = true
	fmt.Print(isOk)
	fmt.Printf("name:%s", name)
	fmt.Println(age)
	fmt.Print(s1)
}
