package main

import (
	"fmt"
	"reflect"
)

type cat struct {
	name string
}

func reflectType(x interface{})  {
	v := reflect.TypeOf(x)
	//fmt.Printf("type:%v\n",v)
	fmt.Printf("type:%v,kind:%v\n",v.Name(),v.Kind())

}

func main() {
	var a float32 = 3.14
	reflectType(a)
	var b int64 = 100
	reflectType(b)
	var c = cat{name:"jack"}
	reflectType(c)
}