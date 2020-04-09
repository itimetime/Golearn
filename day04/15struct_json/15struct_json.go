package main

import (
	"encoding/json"
	"fmt"
)

// json 的序列化和反序列化

type person struct {
	Name string `json:"name",db:"name"`
	Age  int `json:"age"`
}


func main() {
	p1 := person{
		Name: "ck",
		Age:  20,
	}
	b, err := json.Marshal(p1) //序列化
	if err != nil {
		fmt.Printf("error!%v", err)
		return
	}
	fmt.Printf("%v", b) //必须要转成string

	fmt.Printf("%v", string(b))

	//反序列化
	str := `{"name":"ckk","age":22}`
	var p2 person
	json.Unmarshal([]byte(str),&p2)
	fmt.Printf("%#v",p2)
}
