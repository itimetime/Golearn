package main

import (
	"fmt"
	"strings"
)

func makeSuffixFunc(suffix string) func(string) string{
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}else{
			return name
		}
	}

}

func main()  {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test"))
	fmt.Println(jpgFunc("test.jpg"))
	fmt.Println(txtFunc("txt.jpg"))
	fmt.Println(txtFunc("txt.txt"))
}