package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func request()  {
	data := url.Values{}
	data.Set("name","ck")
	data.Set("age","23")
	data.Encode() //如果map中有进行转义
	http.NewRequest("GET", "http://127.0.0.1:9090/xxx/",nil)

}

func main() {
	resp, err :=http.Get("http://127.0.0.1:9090/mian/")
	if err != nil {
		fmt.Println("err!,",err)
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Body)
	fmt.Println(resp.Request)
	fmt.Println(resp.Status)
	fmt.Println(string(bytes))
	resp2, err :=http.Get("http://127.0.0.1:9090/xxx/?name=ck&age=19")
	fmt.Println(resp2.Body)


}
