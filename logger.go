package log

import (
	"strconv"
	"strings"
)

const (
	DebugLevel Level = 2 << iota
	InfoLevel
	ErrorLevel
)

type L interface {
	Output(depth int, level Level, format string, args ...interface{}) error
	SetLevel(level string)
	SetName(name string)
}

type Level int

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case ErrorLevel:
		return "ERROR"
	default:
		return strconv.Itoa(int(l))
	}
}

func ParseLevel(l string) Level {
	switch strings.ToUpper(l) {
	case "DEBUG", "D", "0", "2":
		return DebugLevel
	case "INFO", "I", "4":
		return InfoLevel
	case "ERROR", "ERR", "E", "8":
		return ErrorLevel
	default:
		return InfoLevel
	}
}
