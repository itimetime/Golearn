package main

import "fmt"


func patition(arr *[5]int, left,right int)  {
	if right <= left{
		return
	}
	low := left
	high := right
	base := arr[right]
	for ;left < right;{
		for arr[left] <= base && right > left{
			left ++
		}
		for arr[right] >= base && right > left{
			right --
		}
		if left < right {
			temp := arr[left]
			arr[left] = arr[right]
			arr[right] = temp
		}

	temp := arr[left]
	arr[left] = arr[high]
	arr[high] = temp
	fmt.Println(arr)
	patition(arr,low,left - 1)
	patition(arr,left + 1,high)
	}
}


func main()  {
	arr := [5]int{1,23,3,4896,4}
	left := 0
	right := len(arr) - 1
	patition(&arr,left,right)
	fmt.Println(arr)
}


