package main

import (
	"os"
)

//seek随光标插入内容

func seek()  {
	fileObj, _ := os.OpenFile("./h.text",os.O_RDWR, 0644)
	fileObj.Seek(2,0)
	var s []byte
	s = []byte{'c'}
	fileObj.Write(s)
	//var ret [1]byte
	//n, err := fileObj.Read(ret[:])
	//if err != nil{
	//	fmt.Printf("read error!%s\n",err)
	//}
	defer fileObj.Close()
	//fmt.Println(string(ret[:n]))

}

func main() {
	seek()
}