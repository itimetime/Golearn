package main

import "fmt"

func main(){
	//跳出两次for循环 用flag比较好
	for i := 0; i< 10; i++{
		for j := 0; j < 10; j++{
			if j == 2{
				goto breakTag
			}
		}
	}
	breakTag:
		fmt.Print("break")
}