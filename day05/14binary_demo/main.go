package main

import "fmt"

//二进制实际用途
//三个二进制位表示事件
/// 1 1 1 代表吃饭 睡觉 打豆豆

const (
	eat    int = 4
	sleep  int = 2
	doudou int = 1
)

func f(arg int) {
	fmt.Printf("%b\n", arg)
}

func main() {
	f(eat|doudou)//101
	f(eat|sleep|doudou) //111
}
