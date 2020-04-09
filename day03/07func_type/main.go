package main

import "fmt"

func f1() {
	fmt.Println("hello")
}

func f2() int {
	return 10
}

func f3(x func() int) {
	ret := x()
	fmt.Println(ret)

}

func f4(x, y int) int {
	return x + y
}

//函数可以作为返回值

func f5(x func() int) func(int,int) int{
	ret := func(a,b int) int {
		return a + b
	}
	return ret
}


func main() {
	a := f1
	b := f2
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	f3(f2)
	f6 := f5(f2)
	fmt.Println(f6(4,5))
}
