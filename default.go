package log

import "os"

var standard = Standard()

const depth = 2

func Set(logger L) {
	standard = logger
}

func Config(level, name string, caller, showTime bool) {
	standard.Config(ParseLevel(level), name, caller, showTime)
}

func Debugf(format string, args ...interface{}) {
	_ = standard.Output(depth, DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	_ = standard.Output(depth, InfoLevel, format, args...)
}

func Errorf(format string, args ...interface{}) {
	_ = standard.Output(depth, ErrorLevel, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	_ = standard.Output(depth, ErrorLevel, format, args...)
	os.Exit(1)
}
