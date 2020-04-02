package log

import "os"

var std = New()

func Set(l L) {
	std = l
}

func Config(level, name string, caller, showTime bool) {
	std.Config(ParseLevel(level), name, caller, showTime)
}

func Debugf(format string, args ...interface{}) {
	_ = std.Output(2, DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	_ = std.Output(2, InfoLevel, format, args...)
}

func Errorf(format string, args ...interface{}) {
	_ = std.Output(2, ErrorLevel, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	_ = std.Output(2, ErrorLevel, format, args...)
	os.Exit(1)
}
