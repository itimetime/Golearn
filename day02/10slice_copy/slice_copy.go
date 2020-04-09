package main

import "fmt"

func main() {
	a1 := []int{1, 3,2, 4, 5, 6}
	fmt.Println(len(a1),cap(a1))
	a2 := a1
	a3 := make([]int, 3, 3)
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	//go中没有删除元素的方法 需要append进行拼接
	//eg：删除a1中的3
	fmt.Println("%p\n",&a1[0])
	a4 := append(a1[:1], a1[4:]...)
	//相当于把后面的a1[4:] 前进 覆盖中间的值
	fmt.Println("%p\n",&a1[0])
	//切片不保存具体的值
	//切片对应一个底层数组
	//底层数组都是占用一块连续的内存
	fmt.Println(a4)
	fmt.Println(a1)
}
