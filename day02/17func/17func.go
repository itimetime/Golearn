package main

import "fmt"

//函数
//函数的定义
//函数时一个代码的封装
//

func sum(x int, y int) (ret int) {
	return x + y
}

//没有返回值
func f1(x int, y int) {
	fmt.Printf("%d", x+y)
}

//没有参数没有返回值
func f2() {
	fmt.Println("f3")
}

//没有参数担忧返回值
func f3() int {
	return 3
}

//返回值可以命名也可以不命名
//命名的返回值相当于在函数中声明了一个变量

func f4(x int, y int) (ret int) {
	ret = x + y
	return //ret 可以不写 写也可以
}

//多个返回值
func f5() (int, string) {
	return 1, "6"
}

//参数类型简写 此时x，y均为int类型,m n为string类型
func f6(x, y int, m , n string) int {
	return x+y
}
//可变长参数
//必须放在函数参数的最后
func f7(x string, y ...int)  {
	fmt.Println(x)
	fmt.Println(y)
}
//Go语言中函数没有默认参数这个概念

func main() {
	ans := sum(10, 20)
	fmt.Println(ans)
	f7("你好呀",78,55)
}
