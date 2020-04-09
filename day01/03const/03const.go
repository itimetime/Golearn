package main

import "fmt"

const pi = 3.1415926535
const (
	notFound = 404
	statusOk = 200
)
const (
	n1 = 100
	n2
	n3
)
const (
	a1 = iota //0
	a2 = iota //1
	a3
	_ //匿名变量
	a4
)

const (
	b1 = iota //0
	b2 = 100  //1
	b3 = iota
	_  //匿名变量
	b4
)
const (
	c1, c2 = iota + 1, iota + 2
	c3, c4 = iota + 1, iota + 3
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(n3)
	fmt.Println(a1, a2, a3, a4)
	fmt.Println(b1, b2, b3, b4)
	fmt.Println(c1, c2, c3, c4)

}
