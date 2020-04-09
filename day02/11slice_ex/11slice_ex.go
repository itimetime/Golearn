package main

import (
	"fmt"
	"sort"
)

func main(){
	var a = make([]int, 5, 10)
	fmt.Println(a)
	for i := 0; i< 10;i++{
		a = append(a, i)
		print(i)
	}
	fmt.Println(a)

	res := []int{1, 8, 5, 2, 6}
	sort.Ints(res)
	fmt.Println(res)

}
