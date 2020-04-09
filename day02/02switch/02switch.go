package main

import "fmt"

func main() {
	//基本类型
	n := 5
	switch n {
	case 1:
		fmt.Print(1)
	case 2:
		fmt.Print(2)
	default:
		fmt.Print("error")
	}
	//变种1
	i := 10
	switch i {
	case 1, 3, 5, 7, 9:
		fmt.Print("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Print("偶数")
	default:
		fmt.Print("error")
	}
	//变种2
	j := 5

	switch {
	case j <= 5:
		fmt.Print("<=5")
	default:
		fmt.Print(">=5")

	}

}
