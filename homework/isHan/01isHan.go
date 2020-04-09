package main

import (
	"fmt"
	"unicode"
)

func main() {
	s1 := "曹坤说，good moning 王月"
	n := 0

	for _, i := range (s1) {
		if unicode.Is(unicode.Scripts["Han"], i) {
			fmt.Printf("%c是汉字\n",i)
			n += 1
		}

	}
	fmt.Println(n)
}
