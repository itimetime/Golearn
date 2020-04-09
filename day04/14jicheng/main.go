package main

import "fmt"

//结构题模拟实现其他语言的基础

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s会动", a.name)
}

//狗类

type dog struct {
	feet uint8
	animal //animal使用的方法dog也有了
}

//给狗实现一个汪汪汪的方法

func (d dog) wang() {
	fmt.Printf("%s在叫：汪汪汪", d.name)
}

func main() {
	d1 := dog{
		feet:   4,
		animal: animal{name: "rourou"},
	}
	d1.wang()
	fmt.Println()
	d1.move()
}
