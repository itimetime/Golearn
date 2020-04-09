package main

import "fmt"

//defer 延迟执行 延迟到函数返回的时候进行执行 栈 先进后出
func deferDemo()  {
	fmt.Println("start")
	defer fmt.Println("hahaha")
	defer fmt.Println("ccccc")
	fmt.Println("end")
}

func main(){
	deferDemo()
}

