package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//打开文件


func reader() {
	fileObj, err := os.Open("./file.go")
	if err != nil{
		fmt.Printf("Open file err! message:%s\n",err)
		return
	}
	//记得关闭文件
	defer fileObj.Close()
	//读文件
	//写法1
	//var tmp1 = make([]byte, 128)
	//fileObj.Read(tmp1)
	//写法2
	var tmp2 [128]byte

	for {
		n , err2 :=fileObj.Read(tmp2[:])
		//if n == 0{
		//	return
		//}
		if err2!= nil{
			if err == io.EOF{
				fmt.Printf("读完了")
				return
			}
			fmt.Printf("Read file err! message:%s\n",err)
			return
		}
		//打印出度的数据
		fmt.Printf("读了%n个字节\n",n)
		fmt.Println(string(tmp2[:n]))
	}

}

//利用bufio这个包
func readFromFileBufio()  {
	fileObj, err := os.Open("./file.go")
	if err != nil{
		fmt.Printf("Open file err! message:%s\n",err)
		return
	}
	//记得关闭文件
	defer fileObj.Close()
	//创建一个用来从文件中读取的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF{
			return
		}
		if err != nil{
			fmt.Printf("read line failed, err:%v\n",err)
			return
		}
		fmt.Println(line)
	}

}

func readFromFileByIoutil()  {
	ret , err := ioutil.ReadFile("./file.go")
	if err != nil{
		fmt.Printf("read line failed, err:%v\n",err)
		return
	}
	fmt.Println(ret)
}

func main() {
	//readFromFileBufio()
	readFromFileBufio()
}
