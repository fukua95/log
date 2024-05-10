package log

import (
	"fmt"
	"log"
	"os"
)

var (
	l = NewLogger()
)

const (
	calldepth = 2
)

func header(lvl, msg string) string {
	return fmt.Sprintf("%s: %s", lvl, msg)
}

type Logger struct {
	*log.Logger
	debug bool
}

func NewLogger() *Logger {
	l := &Logger{
		Logger: log.New(os.Stderr, "", 0),
		debug:  false,
	}
	l.EnableTimestamps()
	return l
}

func (l *Logger) EnableTimestamps() {
	l.SetFlags(l.Flags() | log.Ldate | log.Ltime)
}

func (l *Logger) EnableDebug() {
	l.debug = true
}

func Debug(v ...interface{}) {
	if l.debug {
		_ = l.Output(calldepth, header("DEBUG", fmt.Sprint(v...)))
	}
}

func Debugf(format string, v ...interface{}) {
	if l.debug {
		_ = l.Output(calldepth, header("DEBUG", fmt.Sprintf(format, v...)))
	}
}

func Info(v ...interface{}) {
	_ = l.Output(calldepth, header("INFO", fmt.Sprint(v...)))
}

func Infof(format string, v ...interface{}) {
	_ = l.Output(calldepth, header("INFO", fmt.Sprintf(format, v...)))
}

func Error(v ...interface{}) {
	_ = l.Output(calldepth, header("ERROR", fmt.Sprint(v...)))
}

func Errorf(format string, v ...interface{}) {
	_ = l.Output(calldepth, header("ERROR", fmt.Sprintf(format, v...)))
}

func Warning(v ...interface{}) {
	_ = l.Output(calldepth, header("WARN", fmt.Sprint(v...)))
}

func Warningf(format string, v ...interface{}) {
	_ = l.Output(calldepth, header("WARN", fmt.Sprintf(format, v...)))
}

func Fatal(v ...interface{}) {
	_ = l.Output(calldepth, header("FATAL", fmt.Sprint(v...)))
}

func Fatalf(format string, v ...interface{}) {
	_ = l.Output(calldepth, header("FATAL", fmt.Sprintf(format, v...)))
}

func Panic(v ...interface{}) {
	l.Logger.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	l.Logger.Panicf(format, v...)
}
