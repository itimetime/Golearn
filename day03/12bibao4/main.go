package main

import "fmt"

func calc(base int) (func(int) int,func(int) int){
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main()  {
	f1, f2 := calc(10)
	fmt.Println(f1(1),f2(2))
	fmt.Println(f1(3),f2(4))
	fmt.Println(f1(5),f2(6))

	f3, f4 := calc(100)
	fmt.Println(f3(1),f4(2))
	fmt.Println(f4(3),f3(4))
	fmt.Println(f4(5),f4(6))
}
