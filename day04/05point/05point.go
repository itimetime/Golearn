package main

import "fmt"

func main()  {
	var a int
	a = 100
	b := &a
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	fmt.Printf("%p\n",b)
	fmt.Printf("%p\n",&a)



}
