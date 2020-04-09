package main

import "fmt"

func main() {
	//基本格式 此时i的作用域在for的内部
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	//变种1 i可以在后面使用 且循环后变为10
	i := 5
	for ; i < 10; i++ {
		fmt.Print(i)
	}

	//变种2
	j := 5
	for ; j < 1; {
		j++
	}

	//for range 循环

	s := "所有的苦难都会过去的"

	for i, j := range s{
		fmt.Printf("%d, %c\n", i, j)
	}


}
