package go_logging

import (
	"fmt"
	"runtime"
	"time"
	"strings"
)

type ILogger interface {
	SetLevel(level int)
	GetLevel() int
	Info(a ...interface{})
	Warn(a ...interface{})
	Error(a ...interface{})
	Fatal(a ...interface{})
}

type Logger struct {
	level int
}

func (logger *Logger) SetLevel(level int) {
	logger.level = level
}

func (logger *Logger) GetLevel() int {
	return logger.level
}

func getLogContextString() string {
	pc, filename, line, ok := runtime.Caller(2)
	if ok {
		fn := runtime.FuncForPC(pc).Name()
		return fmt.Sprintf("%s[%s:%d]", fn, filename, line)
	} else {
		return fmt.Sprintf("<UNKOWN>[<UNKNOWN>:<UNKNOWN>]")
	}
}

func getShortLogContextString() string {
	pc, filename, line, ok := runtime.Caller(2)
	if ok {
		fn := runtime.FuncForPC(pc).Name()
		return fmt.Sprintf("%s[%s:%d]",
			fn[strings.LastIndex(fn, "/")+1:],
			filename[strings.LastIndex(filename, "/")+1:],
			line)
	} else {
		return fmt.Sprintf("<UNKOWN>[<UNKNOWN>:<UNKNOWN>]")
	}
}

func getLogTimeString() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

func (logger *Logger) Info(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	//long_msg := fmt.Sprintf("%s [info] in %s %s", getLogTimeString(), getLogContextString(), msg)
	short_msg := fmt.Sprintf("[info] %s %s", getShortLogContextString(), msg)
	println(short_msg)
}

func (logger *Logger) Warn(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	//long_msg := fmt.Sprintf("%s [warn] in %s %s", getLogTimeString(), getLogContextString(), msg)
	short_msg := fmt.Sprintf("[warn] %s %s", getShortLogContextString(), msg)
	println(short_msg)
}

func (logger *Logger) Error(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	//long_msg := fmt.Sprintf("%s [error] in %s %s", getLogTimeString(), getLogContextString(), msg)
	short_msg := fmt.Sprintf("[error] %s %s", getShortLogContextString(), msg)
	println(short_msg)
}

func (logger *Logger) Fatal(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	//long_msg := fmt.Sprintf("%s [fatal] in %s %s", getLogTimeString(), getLogContextString(), msg)
	short_msg := fmt.Sprintf("[fatal] %s %s", getShortLogContextString(), msg)
	println(short_msg)
}
