package main

import "fmt"

func main(){
	var i1  = 101
	fmt.Printf("%d\n",i1)
	fmt.Printf("%b\n",i1)
	fmt.Printf("%o\n",i1)
	fmt.Printf("%x\n",i1)

	//八进制
	i2 := 077
	fmt.Print(i2)
	//十六进制
	i3 := 0x123456
	fmt.Printf("%d\n",i3)
	fmt.Print(i3)

	//声明int8类型的变量,否则为默认
	i4 := int8(9)
	fmt.Printf("%T\n",i4)

}
