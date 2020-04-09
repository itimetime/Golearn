package main

import "fmt"

type myString string


var name myString


func (s myString)hello()  {
	fmt.Printf("%s:Hello!",s)
}

func main()  {
	haHa := myString("lalala")
	name = "ck"
	name.hello()
	haHa.hello()
}