package main

import "fmt"

//结构体是值类型


type person struct {
	name, gender string
}

func f(x person)  {
	x.gender = "stonger man"
}

func f2(x *person)  {
	(*x).gender = "superman"  //可以直接x.gender = "superman"
}
func main()  {
	var p person
	p.name = "ck"
	p.gender = "boy"
	f(p)
	fmt.Println(p)
	f2(&p)
	fmt.Println(p)
	//结构体指针1
	var p2  = new(person)  //生成的指针
	p2.gender = "ff"
	fmt.Printf("%T\n",p2)
	fmt.Printf("%p\n",p2)  //p2 保存的值的内存地址
	fmt.Printf("%p\n",&p2) //p2的内存地址
	//结构体指针2 1.key-value初始化
	var p3  = &person{  //拿地址
		name:   "ckk",
		gender: "boy",
	}
	fmt.Printf("%#v\n",p3)
	//值的顺序必须和 定义的结构体一致
	p4 := person{
		name:   "ccccc",
		gender: "boy",
	}
	fmt.Printf("%#v\n",p4)


}
