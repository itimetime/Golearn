package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//打开文件写内容

func write1()  {
	fileObj , err := os.OpenFile("./xx.text", os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0644)
	if err != nil{
		fmt.Printf("open file error %v",err)
	}
	//write
	fileObj.Write([]byte("???要重新复习\n"))
	fileObj.WriteString("直接写入文件")
	fileObj.Close()
}

func write2()  {
	fileObj , err := os.OpenFile("./xx1.text", os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0644)
	if err != nil{
		fmt.Printf("open file error %v",err)
	}
	defer fileObj.Close()
	writer := bufio.NewWriter(fileObj)
	writer.WriteString("hahhahah")
	writer.Flush()
}

func write3()  {
	str := "hello ssshh"
	err := ioutil.WriteFile("./xx3.text",[]byte(str),0666)
	if err != nil{
		fmt.Printf("write file error %v",err)
	}


}

func main() {
	write1()
	write2()
	write3()
}