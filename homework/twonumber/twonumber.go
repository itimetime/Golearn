package main

import "fmt"

func main() {
	res := [5]int{1, 3, 5, 7, 8}
	for i, j := range res {
		for k := i + 1; k < len(res); k++ {
			if j+res[k] == 8 {
				fmt.Println(i,k)
			}else if j+res[k] < 8{
				continue
			}else {
				break
			}

		}
	}

}
