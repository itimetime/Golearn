package main

import "fmt"

// pointer
func main() {
	//1. &：取地址
	//2. *：根据地址取值
	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	n2 := *p
	fmt.Println(n2)

	var a1 *int
	fmt.Println(a1)
	//new 函数申请一个内存地址
	var a2 = new(int)
	fmt.Println(*a2)
	fmt.Println(a2)
	*a2 = 100
	fmt.Println(*a2)
}
