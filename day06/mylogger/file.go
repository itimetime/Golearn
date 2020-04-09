package mylogger

import (
	"fmt"
	"time"
)

type fileLogger struct {
	level    LogLevel
	filePath string //文件路径
	fileName string //保存的文件名称
	maxFile  int64
}

func NewFileLogger(logLever, fb, fn string, maxSize int64) (*fileLogger, error) {
	level, err := parseLogLevel(logLever)
	return &fileLogger{level: level,
		filePath: fb,
		fileName: fn,
		maxFile:  maxSize,
	}, err

}

func (l fileLogger) enable(level LogLevel) bool {
	if level >= l.level {
		return true
	} else {
		return false
	}
}

func fileWriteLog(lv LogLevel, format string, a ...interface{})  {
	msg := fmt.Sprintf(format ,a...)
	funcName, fileName, line := getInfo(2)
	now := time.Now().Format("2006/01/02 15:04:05")
	fmt.Printf("[%s] [%s] [%s:%s:%d]%s\n", now, getLogString(lv),funcName, fileName,line,msg)
}


func (l fileLogger) Debug(format string, a ...interface{}) {
	if l.enable(DEBUG) {
		//fmt.Fprint(os.Stdout,msg)
		log(DEBUG,format, a...)
	}

}

func (l fileLogger) Info(format string, a ...interface{}) {
	if l.enable(INFO) {
		log(INFO,format, a...)
	}

}

func (l fileLogger) Warning(format string, a ...interface{}) {
	if l.enable(WARNING) {
		log(WARNING,format, a...)
	}

}
func (l fileLogger) Error(format string, a ...interface{}) {
	if l.enable(ERROR) {
		log(ERROR,format, a...)
	}

}


func (l fileLogger) Fatal(format string, a ...interface{}) {
	if l.enable(FATAL) {
		log(FATAL,format, a...)
	}

}
