package main

import "fmt"

//自定义类型和类型别名

type myInt int
type yourInt = int  //类型别名


func main()  {
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n",n)

	var m yourInt
	m = 100
	fmt.Println(m)
	fmt.Printf("%T\n",m)

	var c rune
	c = '中'
	fmt.Println(c)
	fmt.Printf("%T",c)


}
