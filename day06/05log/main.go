package main

//测试我们的日志库
import (
	"golearnProject/day06/mylogger"
)


func main(){
	log, _ := mylogger.NewLog("2020年时间记录器","DEBUg")
	err := "loading failed!"
	log.Debug("test")
	log.Info("test")
	log.Warning("test")
	log.Error("test,err:%s",err)
	log.Fatal("test")

}
