package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func Standard() L {
	return &logger{
		l:      log.New(os.Stderr, "", 0),
		level:  DebugLevel,
		name:   "shu.run",
		caller: true,
		time:   true,
	}
}

type logger struct {
	l      *log.Logger
	name   string
	caller bool
	time   bool
	level  Level
}

func (l *logger) Config(level Level, name string, caller, showTime bool) {
	l.caller = caller
	l.time = showTime
	l.level = level
	l.name = name
}

func (l *logger) Output(depth int, level Level, format string, args ...interface{}) error {
	if level >= l.level {

		prefix := fmt.Sprintf("%s.%s ", l.name, level.String()[:4])
		if l.time {
			prefix += fmt.Sprintf("[%s] ", time.Now().In(time.FixedZone("CST", 8*3600)).Format("2006/01/02 15:04:05"))
		}

		if l.caller {
			_, f, l, _ := runtime.Caller(depth)
			_, n := filepath.Split(f)
			prefix += fmt.Sprintf("(%s:%d) ", n, l)
		}

		return l.l.Output(depth, prefix+fmt.Sprintf(format, args...))
	}
	return nil
}

var _ L = (*logger)(nil)
