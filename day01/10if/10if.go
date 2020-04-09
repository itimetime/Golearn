package main

import "fmt"

func main() {
	age := 19
	if age > 35 {
		fmt.Println("中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else{
		fmt.Println("未成年")
	}

	//下面的age只在内部可见
	if age:= 19; age > 35 {
		fmt.Println("中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else{
		fmt.Println("未成年")
	}
}
