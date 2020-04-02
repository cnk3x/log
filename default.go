package log

import "os"

var std = New()

func Set(l L) {
	std = l
}

func SetLevel(level string) {
	std.SetLevel(level)
}

func Name(name string) {
	std.SetName(name)
}

func Debugf(format string, args ...interface{}) {
	_ = std.Output(3, DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	_ = std.Output(3, InfoLevel, format, args...)
}

func Errorf(format string, args ...interface{}) {
	_ = std.Output(3, ErrorLevel, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	_ = std.Output(3, ErrorLevel, format, args...)
	os.Exit(1)
}
