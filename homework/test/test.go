package main

import "fmt"

func test(arr *[5]int)  {
	arr[4] = 1000
}
func main()  {
	arr := [5]int{1,2,3,4,6}
	test(&arr)
	fmt.Println(arr)

}
