package log

import (
	"fmt"
	"github.com/op/go-logging"
)

var module string = "keigo"
var logger = logging.MustGetLogger(module)

func init() {
	DisableLog()
}

func VerboseLog() {
	logging.SetLevel(logging.DEBUG, module)
}

func DisableLog() {
	logging.SetLevel(logging.CRITICAL, module)
}

func UseLogger(newLogger *logging.Logger) {
	if newLogger == nil {
		panic("logger is nil")
	}
	logger = newLogger
}

func Debugf(format string, params ...interface{})    { logger.Debug(format, params...) }
func Infof(format string, params ...interface{})     { logger.Info(format, params...) }
func Warnf(format string, params ...interface{})     { logger.Warning(format, params...) }
func Errorf(format string, params ...interface{})    { logger.Error(format, params...) }
func Criticalf(format string, params ...interface{}) { logger.Critical(format, params...) }
func Panicf(format string, params ...interface{})    { logger.Panicf(format, params...) }
func Fatalf(format string, params ...interface{})    { logger.Fatalf(format, params...) }

func Debug(v ...interface{})    { logger.Debug("%s", fmt.Sprint(v...)) }
func Info(v ...interface{})     { logger.Info("%s", fmt.Sprint(v...)) }
func Warn(v ...interface{})     { logger.Warning("%s", fmt.Sprint(v...)) }
func Error(v ...interface{})    { logger.Error("%s", fmt.Sprint(v...)) }
func Critical(v ...interface{}) { logger.Critical("%s", fmt.Sprint(v...)) }
func Panic(v ...interface{})    { logger.Panic(v...) }
func Fatal(v ...interface{})    { logger.Fatal(v...) }
