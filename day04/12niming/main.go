package main

import (
	"fmt"
)

//匿名字段
//适用与字段比较少，也比较简单的字段

type person struct {
	string
	int
}

func main()  {
	p1 := person{"ck",23}
	fmt.Println(p1.string)
	fmt.Println(p1.int)
	//只使用一次的结构体
	var a = struct {
		x int
		y int
	}{3,4}
	fmt.Println(a)
}