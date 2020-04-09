package main

import "fmt"

//闭包是什么
//特点：包含外部作用域的一个变量

func adder(x int) func(int) int{
	return func(y int) int {
		x += y
		return x
	}
}

func main()  {
	ret := adder(100)
	ret2 := ret(200)
	fmt.Println(ret2)
}