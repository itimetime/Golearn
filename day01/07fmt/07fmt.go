package main

import "fmt"

func main()  {

	//fmt占位符  \查看内存地址
	n := 100
	fmt.Printf("%d\n",n)
	fmt.Printf("%v\n",n)
	fmt.Printf("%b\n",n)
	fmt.Printf("%c\n",n)
	fmt.Printf("%o\n",n)
	fmt.Printf("%x\n",n)
	var s = "Colin"
	fmt.Printf("%s\n",s)
	fmt.Printf("%v\n",s)
	fmt.Printf("%#v\n",s)
}
