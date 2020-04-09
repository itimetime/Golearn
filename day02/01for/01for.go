package main

import "fmt"

func main()  {
	//break
	for i := 0; i < 10; i++{
		if i == 5{
			break
		}
		fmt.Print(i)
	}
	//continue
	fmt.Println()
	for i := 0; i< 10; i++{
		if i == 5{
			continue
		}
		fmt.Print(i)
	}


}
