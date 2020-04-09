package main

import (
	"fmt"
	"path"
	"runtime"
)

func f1(){
	caller, file, line, ok := runtime.Caller(0)
	if !ok{
		fmt.Println("runtime.Caller() failed!")
		return
	}
	funcName := runtime.FuncForPC(caller).Name()
	fmt.Println(funcName)
	fmt.Println(path.Base(file))
	fmt.Println(line)
}

func main() {
	f1()
}