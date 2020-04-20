package main

import (
	"fmt"
)

var a []int
var b chan int //需要指定通道中元素的类型

func main() {
	b = make(chan int) //后面可加容量
	go func() {
		x := <-b
		fmt.Println("取到通道中的数值，", x)
	}()

	b <- 10 //无缓冲区 不能放数据 必须等着取出才能放
	fmt.Println("发送到通道中了")

	b = make(chan int, 16) //带缓冲区的通道
	b <- 10
	fmt.Println(b)
	close(b) //关闭通道
}
