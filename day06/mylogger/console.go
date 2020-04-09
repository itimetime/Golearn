package mylogger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

/*
日志库
需求：   1 支持往不同的地方输出日志
		2日志分级别
			1.Debug
			2.Info
			3.Waring
			4.Error
			5.Fatal
		3日志支持开关控制
		4 时间/行号/文件名/日志级别/日志信息
        5 日志文件要切割
*/

// Logger构造函数
func NewLog(name string, logLever string) (ConsoleLogger, error) {
	level, err := parseLogLevel(logLever)
	return ConsoleLogger{name: name, level: level}, err
}

func (l ConsoleLogger) enable(level LogLevel) bool {
	if level >= l.level {
		return true
	} else {
		return false
	}
}

func log(lv LogLevel, format string, a ...interface{})  {
	msg := fmt.Sprintf(format ,a...)
	funcName, fileName, line := getInfo(2)
	now := time.Now().Format("2006/01/02 15:04:05")
	fmt.Printf("[%s] [%s] [%s:%s:%d]%s\n", now, getLogString(lv),funcName, fileName,line,msg)
}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}


func (l ConsoleLogger) Debug(format string, a ...interface{}) {
	if l.enable(DEBUG) {
		//fmt.Fprint(os.Stdout,msg)
		log(DEBUG,format, a...)
	}

}

func (l ConsoleLogger) Info(format string, a ...interface{}) {
	if l.enable(INFO) {
		log(INFO,format, a...)
	}

}

func (l ConsoleLogger) Warning(format string, a ...interface{}) {
	if l.enable(WARNING) {
		log(WARNING,format, a...)
	}

}
func (l ConsoleLogger) Error(format string, a ...interface{}) {
	if l.enable(ERROR) {
		log(ERROR,format, a...)
	}

}


func (l ConsoleLogger) Fatal(format string, a ...interface{}) {
	if l.enable(FATAL) {
		log(FATAL,format, a...)
	}

}

func getInfo(n int) (funcName, fileName string, line int) {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("runtime.Caller() failed!")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	s := strings.Split(funcName, "/")
	funcName = s[len(s)-1]
	fileName = path.Base(file)
	return
}
