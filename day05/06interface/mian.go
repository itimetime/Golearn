package main

import "fmt"

// 使用值接收者和指针接受者的区别
//接口的底层逻辑  分为两部分 存储类型和值（动态类型和动态值）
type animal interface {
	move()
	eat(food string)
	//直接写成 eat(string)
}

type cat struct {
	name string
	legs int8
}


type chicken struct {
	name string
	legs int8
}

func (c *cat)move()  {
	fmt.Printf("%s走路\n",c.name)
}

func (c *cat) eat (food string)  {
	fmt.Printf("%s吃%s\n",c.name,food)
}

func (c chicken)move()  {
	fmt.Printf("%s走路\n",c.name)
}

func (c chicken) eat (food string)  {
	fmt.Printf("%s吃%s\n",c.name,food)
}

func feed(a animal, food string) {
	a.eat(food)
	a.move()
}

func main()  {
	var a1 animal
	c1 := cat{"tom",4}
	c2 := &cat{"假老练",4}
	//a1 = c1  //这样写会报错
	a1 = &c1
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)

}