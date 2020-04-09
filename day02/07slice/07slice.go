package main

import "fmt"

//切片

func main() {
	var s1 []int //切片初始化
	var s2 []string
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	s1 = []int{1, 2, 3}
	s2 = []string{"a", "b", "c"}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	fmt.Println(len(s1),cap(s1))
	fmt.Println(len(s2),cap(s2))
	a1 := [...]int{1,3,5,7,9,11,13}
	s3 := a1[0:4] //切片方法和python一样
	s4 := a1[2:3]
	s5 := s3[0:2]
	fmt.Println(s3)
	//切片是引用，底层数组改变 切片也会改变
	fmt.Println(cap(s3)) //切片的容量指的是底层 数组的容量 内存只能向右扩展
	fmt.Println(cap(s4)) //切片的容量指的是底层 数组的容量
	fmt.Println(cap(s5)) //切片的容量指的是底层 数组的容量
	s5[0] = 10000
	//切边改变底层数组也会改变
	fmt.Println(a1)

	//make函数建立切片
}
