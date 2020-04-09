package main

import "fmt"

//go 语言中return分为两步
//1 返回值赋值
//2 return

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4()(x int)  {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func f5() (x int){
	defer func(x int) int {
		x++
		return x
	}(x)
	return 5
}

func f6() (x int){
	defer func(x *int)  {
		(*x)++
	}(&x)
	return 5
}


func main()  {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
	fmt.Println(f6())
}
