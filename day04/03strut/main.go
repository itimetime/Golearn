package main

import "fmt"

type person struct {
	name string
	age int
	gender string
	hobby []string
}

func main()  {
	var p person
	p.name = "ck"
	p.age = 23
	p.gender = "boy"
	p.hobby = []string{"game","film"}

	fmt.Println(p)

	var s struct{
		x int
		y string
	}
	s.x = 100
	s.y = "1000"
	fmt.Println(s)

}