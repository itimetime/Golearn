package main

import "fmt"

// make()函数创造切片

func main() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1 = %v len(s1)= %d cap(s1) = %d \n", s1, len(s1), cap(s1))
	s2 := make([]int, 0, 10)
	fmt.Printf("s1 = %v len(s1)= %d cap(s1) = %d \n", s2, len(s2), cap(s2))
	//切片的本质  底层数值的指针，长度，容量。属于引用类型

}
