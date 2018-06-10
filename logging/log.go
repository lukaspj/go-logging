package logging

import (
	"fmt"
	"runtime"
	"time"
	"strings"
	"os"
)

var logger = new(Logger)

const (
	ALL   = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF
)

type ILogger interface {
	SetLevel(level int)
	GetLevel() int
	AddLogOutput(output ILogOutput)
	Info(a ...interface{})
	Warn(a ...interface{})
	Error(a ...interface{})
	Fatal(a ...interface{})
}

type ILogOutput interface {
	Println(message string)
}

type Logger struct {
	level  int
	output []ILogOutput
}

func GetLogger() *Logger {
	return logger
}

func (logger *Logger) AddLogOutput(output ILogOutput) {
	logger.output = append(logger.output, output)
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

func (logger *Logger) Log(level int, format string, a ...interface{}) {
	if logger.level <= level {
		for _, out := range logger.output {
			out.Println(fmt.Sprintf(format, a...))
		}
	}
}

func (logger *Logger) Info(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	//long_msg := fmt.Sprintf("%s [info] in %s %s", getLogTimeString(), getLogContextString(), msg)
	short_msg := fmt.Sprintf("[info] %s %s", getShortLogContextString(), msg)
	logger.Log(INFO, short_msg)
}

func (logger *Logger) Warn(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	//long_msg := fmt.Sprintf("%s [warn] in %s %s", getLogTimeString(), getLogContextString(), msg)
	short_msg := fmt.Sprintf("[warn] %s %s", getShortLogContextString(), msg)
	logger.Log(WARN, short_msg)
}

func (logger *Logger) Error(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	//long_msg := fmt.Sprintf("%s [error] in %s %s", getLogTimeString(), getLogContextString(), msg)
	short_msg := fmt.Sprintf("[error] %s %s", getShortLogContextString(), msg)
	logger.Log(ERROR, short_msg)
}

func (logger *Logger) Fatal(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	//long_msg := fmt.Sprintf("%s [fatal] in %s %s", getLogTimeString(), getLogContextString(), msg)
	short_msg := fmt.Sprintf("[fatal] %s %s", getShortLogContextString(), msg)
	logger.Log(FATAL, short_msg)
}

type StdoutOutput struct{}

func (out *StdoutOutput) Println(message string) {
	println(message)
}

func (logger *Logger) AddStdoutOutput() {
	output := new(StdoutOutput)
	logger.AddLogOutput(output)
}

type FileOutput struct {
	FileName string
	output   chan string
	done     chan bool
}

func (out *FileOutput) Println(message string) {
	out.output <- message
}

func GetFileOutput(fileName string) (output *FileOutput) {
	output = new(FileOutput)
	output.done = make(chan bool)
	output.output = make(chan string)
	go func() {
		f, err := os.Create(fileName)
		if err != nil {
			logger.Error("error %v", err)
			return
		}

		defer f.Close()

		for {
			select {
			case s := <-output.output:
				f.WriteString(s + "\n")
			case <-output.done:
				break
			}
		}
	}()
	return
}
