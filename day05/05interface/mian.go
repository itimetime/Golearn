package main

import "fmt"

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

func (c cat)move()  {
	fmt.Printf("%s走路\n",c.name)
}

func (c cat) eat (food string)  {
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
	var cat1 = cat{"蓝猫",4}
	var chicken1 = chicken{"🐥",2}
	feed(cat1,"猫粮")
	feed(chicken1,"鸡饲料")
	a1 = cat1
	fmt.Printf("%T\n",a1)
	fmt.Println(a1)
	a1 = chicken1
	fmt.Printf("%T\n",a1)
	fmt.Println(a1)


}