package main

import "fmt"

//panic 和 recover

func funcA(){
	fmt.Println("A")
}

func funcB(){
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("处理 error")
	}()
	panic("error!!!!!!!!!")
	fmt.Println("B")
}

func funcC(){
	fmt.Println("C")
}

func main(){
	funcA()
	funcB()
	funcC()

}


