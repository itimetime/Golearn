package main

import "fmt"

func main()  {
	a := [3]int{1,23,3}
	f1(a) //没有修改a，相当于传递进函数的时副本
	//和python不同 python会进行修改
	fmt.Println(a)

	s1 := []int{1,2,3}
	s2 := s1
	var s3 []int //未开辟新空间
	var s4 = make([]int,0,3) //长度为0,无法存储
	var s5 = make([]int,3)
	copy(s3,s1)
	copy(s4,s1)
	copy(s5,s1)
	s2[1] = 200
	fmt.Println(s2)
	fmt.Println(s1)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)

}

func f1(a [3]int)  {
	a[1] = 100
}