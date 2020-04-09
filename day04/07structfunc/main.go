package main

import (
	"fmt"
)

//构造函数
//当结构体比较小 可以返回person
//当结构体比较大， 可以使用指针


type person struct {
	name string
	age  int
}

func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

func newPerson2(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("ck", 15)
	p2 := newPerson("wy", 13)
	fmt.Println(p1, p2)
}
