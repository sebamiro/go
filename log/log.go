package log

import (
	"fmt"
	"runtime"
	"strings"
)

type logLevel uint8

const (
	TRACE logLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var logLevelStrings = []string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
var logLevelColor = []string{"37", "1;35", "1;32", "1;33", "1;31", "31"}

var currentLogLvl = INFO

func SetLogLevel(l logLevel) {
	if l > FATAL {
		return
	}
	currentLogLvl = l
}

func (l logLevel) String() string {
	return "[" + logLevelStrings[l] + "]"
}

func Tracef(f string, args ...interface{}) {
	logF(TRACE, f, args...)
}

func Debugf(f string, args ...interface{}) {
	logF(DEBUG, f, args...)
}

func Infof(f string, args ...interface{}) {
	logF(INFO, f, args...)
}

func Warnf(f string, args ...interface{}) {
	logF(WARN, f, args...)
}

func Errorf(f string, args ...interface{}) {
	logF(ERROR, f, args...)
}

func Fatalf(f string, args ...interface{}) {
	logF(FATAL, f, args...)
}

func logF(logLvl logLevel, f string, args ...interface{}) {
	if logLvl < currentLogLvl {
		return
	}

	var caller string = "unknownCaller"
	_, file, line, ok := runtime.Caller(2)
	if ok {
		fileName := file[strings.LastIndex(file, "/")+1:]
		caller = fmt.Sprintf("%s:%d", fileName, line)
	}
	fmt.Printf(color(logLvl, fmt.Sprintf("%+7s %s | %s\n", logLvl, caller, f)), args...)
}

func color(logLvl logLevel, s string) string {
	c := logLevelColor[logLvl]
	return fmt.Sprintf("\x1b[%sm%s\x1b[0m", c, s)
}
