package main

import "fmt"

//同一个结构体可以实现多个接口

type mover interface {
	move()
}

type eater interface {
	eat()
}

//type animal interface {
//	move()
//	eat()
//}
//可以直接嵌套

type animal interface {
	mover
	eater
}



type cat struct {
	name string
	legs int8
}
//cat 实现了move接口
func (c *cat)move()  {
	fmt.Printf("%s走路\n",c.name)
}
//cat实现了eat接口
func (c *cat) eat (food string)  {
	fmt.Printf("%s吃%s\n",c.name,food)
}
func main() {

}