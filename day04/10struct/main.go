package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var p1 person
	p1.name = "hahh"
	p1.age = 15

	p2 := person{name: "ffff", age: 78}
	fmt.Println(p2)

	var p3 = person{name: "xxx", age: 15}
	fmt.Println(p3)

	var p4 = person{"xxx",  15}
	fmt.Println(p4)

	p5 := newPerson("ck",23)
	fmt.Println(p5)

	s1 := []int{1,2,3,4}
	s2 := make([]string,3,6)
	fmt.Println(s1,s2)

}

func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}
