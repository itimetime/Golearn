package main

import "fmt"

//方法
//是作用于特定类型的函数
//标识符首字母大写 代表对外部包可见 共有的 可被别的包调用的

//狗的结构体
type dog struct {
	name string
}
//构造函数
func newDog(name string) dog {
	return dog{name: name}
}
//方法是作用域特定类型的函数
//接受者表示的是调用还方法的具体类型的变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪",d.name)
}


type person struct {
	age int
}

func newPerson(age int) person {
	return person{age: age}
}

func (p person) getAge() {
	p.age++
}

//使用指针接收者
func (p *person) getAge2() {
	p.age++
}


func main() {
	d1 := newDog("rourou")
	d1.wang()

	p1 := newPerson(15) //使用值接受方法 不会进行改变
	p1.getAge()
	fmt.Println(p1)

	p2 := newPerson(15) //使用值接受方法 不会进行改变
	p2.getAge2()
	fmt.Println(p2)
}
