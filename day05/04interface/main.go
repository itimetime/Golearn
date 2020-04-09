package main

import "fmt"

//接口实例2
//不管什么车，都能跑
//接口的类型
/*
type 接口名 interface{
    方法一(参数1，参数2)(返回值1，返回值2)
    方法二(参数1，参数2)(返回值1，返回值2)
    ...
}
 */

type car interface {
	run()
}

type falili struct {
	brand string
	speed int
}

type baoshijie struct {
	brand string
	speed int
}

func drive(c car){
	c.run()
}

func (b baoshijie)run()  {
	fmt.Printf("%s,速度%d迈\n",b.brand,b.speed)
}


func (f falili)run()  {
	fmt.Printf("%s,速度%d迈\n",f.brand,f.speed)
}

func main()  {
	var f1 = falili{"法拉利",60}
	var f2 = baoshijie{"保时捷",80}
	drive(f1)
	drive(f2)
}
