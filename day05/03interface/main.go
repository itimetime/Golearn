package main

import "fmt"

//引出接口的实例
//在编程中会遇到以下场景 不关注类型 只关注方法
//接口时一种类型，约束了变量的方法

type speaker interface {
	speak() //只要实现了speak()方法，都能当成speaker()
}

type cat struct {

}

type dog struct {

}


type person struct {

}

func (c cat) speak()  {
	fmt.Println("miaomiaomaio")
}

func (d dog) speak()  {
	fmt.Println("wangwangwang")
}

func (p person) speak()  {
	fmt.Println("yingyingying")
}

func touch(x speaker)  {
	//接受一个参数，传进来谁就喂谁
	//触摸会触发speak方法
	x.speak()
}

func main()  {
	var (
		c1 cat
		d1 dog
		p1 person
	)

	touch(c1)
	touch(d1)
	touch(p1)

	var ss speaker
	ss = c1
	fmt.Println(ss)
	ss = d1
	fmt.Println(ss)
	ss = p1
	fmt.Println(ss)

}