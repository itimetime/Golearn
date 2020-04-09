package main

import (
	"fmt"
)

func f1(f func()) string{
	fmt.Println("this is f1")
	f()
	return "success"
}

func f2(x, y int) {
	fmt.Println("this is f2")
}

//定义一个函数对f2进行包装

func f3(f func(int, int), m, n int) func() {
	tmp := func() {
		fmt.Println("function tmp")
		fmt.Println("开始进行包装")
		f(m,n)
	}
	return tmp
}

func main() {
	ret := f3(f2, 100, 200)
	//闭包等于函数+外部变量
	ret2 := f1(ret)
	fmt.Println(ret2)
}
