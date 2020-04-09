package main

import "fmt"

//数组 必须指定存放位置的类型和容量
//数组的长度是数组类型的一部分
func main() {
	var (
		a1 [3]bool //[true, false, true]
		a2 [4]bool
	)
	fmt.Printf("%T, %T", a1, a2)

	//数组的初始化,默认是0值，bool型为false
	// int float 0
	// 字符串 ""
	fmt.Printf("%v, %v", a1, a2)

	//1.初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)

	//2.初始化方式2 根据初始值自动推算出来数组长度的是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a10)

	//3.初始化方式3 根据索引初始化
	a3 := [5]int{1, 2}
	fmt.Println(a3)
	a4 := [5]int{0: 1, 4: 2}
	fmt.Println(a4)
	//多维数组
	//[[1,2],[3,4],[5,6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a11)
	//多维数组的遍历

	for _, v := range a11 {
		for _, j := range v {
			fmt.Print(j)
		}
		fmt.Println()
	}
	//数组的值类型

	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1,b2)

}
