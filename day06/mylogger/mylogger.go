package mylogger

import (
	"errors"
	"strings"
)

/*
日志库
需求：   1 支持往不同的地方输出日志
		2日志分级别
			1.Debug
			2.Info
			3.Warning
			4.Error
			5.Fatal
		3日志支持开关控制
		4 时间/行号/文件名/日志级别/日志信息
        5 日志文件要切割
 */
type LogLevel uint16

const (
	DEBUG LogLevel = iota
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

//Logger 暴露的要加注释
type ConsoleLogger struct {
	name string
	level LogLevel
}

func parseLogLevel(s string) (LogLevel,error) {
	s = strings.ToUpper(s)
	switch s {
	case "DEBUG":
		return DEBUG, nil
	case "INFO":
		return INFO, nil
	case "TRACE":
		return TRACE, nil
	case "ERROR":
		return ERROR, nil
	case "WARING":
		return WARNING, nil
	case "FATAL":
		return FATAL, nil
	default:
		//err:= fmt.Errorf("无效的日志级别")
		err:= errors.New("无效的日志级别")
		return DEBUG, err
	}
}
