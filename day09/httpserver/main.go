package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request)  {
	//str := "Hello world!"
	b, err := ioutil.ReadFile("./xx.html")
	if err != nil {
		b = []byte("404 Not Found!")
	}
	w.Write(b)
}

func f2(w http.ResponseWriter, r *http.Request)  {
	//str := "Hello world!"
	queryParam := r.URL.Query()
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name,age)
	w.Write([]byte("ok"))
}


func main() {
	http.HandleFunc("/mian/",f1)
	http.HandleFunc("/xxx/",f2)
	http.ListenAndServe("127.0.0.1:9090",nil)

}
