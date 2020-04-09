package main

import "fmt"

//结构体嵌套

type address struct {
	province string
	city string
}


type person struct {
	name string
	age int
	addr address
}

type person1 struct {
	name string
	age int
	address
}


type company struct {
	name string
	addr address
}

func main()  {
	p1 := person{
		name: "ck",
		age:  23,
		addr: address{"山东","菏泽"},
	}
	fmt.Println(p1)
	fmt.Println(p1.addr.city)

	p2 := person1{
		name: "ck",
		age:  23,
		address: address{"山东","菏泽"},
	}
	fmt.Println(p2)
	fmt.Println(p2.city) //匿名嵌套结构体 
}