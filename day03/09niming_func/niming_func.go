package main

import (
	"fmt"
)
//匿名函数一般用在函数内部，因为函数内部是不允许存在变量名字

func main()  {
	a := func(x,y int) {
		fmt.Println(x+y)
	}
	a(3,6)

	func(){
		fmt.Println(666)
	}()
	//后面加括号直接执行
	func(a,b int){
		fmt.Println(a+b)
	}(5,6)
	//后面的括号可以传递变量

}
