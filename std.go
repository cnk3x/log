package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func New() L {
	return &osSTD{
		l:      log.New(os.Stderr, "", 0),
		level:  DebugLevel,
		name:   "shu.run",
		caller: true,
		time:   true,
	}
}

type osSTD struct {
	l      *log.Logger
	name   string
	caller bool
	time   bool
	level  Level
}

func (s *osSTD) Config(level Level, name string, caller, showTime bool) {
	s.caller = caller
	s.time = showTime
	s.level = level
	s.name = name
}

func (s *osSTD) Output(depth int, level Level, format string, args ...interface{}) error {
	if level >= s.level {

		prefix := fmt.Sprintf("%s.%s", s.name, level.String()[:4])
		if s.time {
			prefix += fmt.Sprintf("[%s]", time.Now().In(time.FixedZone("CST", 8*3600)).Format("2006/01/02 15:04:05"))
		}

		if s.caller {
			_, f, l, _ := runtime.Caller(depth)
			_, n := filepath.Split(f)
			prefix += fmt.Sprintf("(%s:%d)", n, l)
		}

		return s.l.Output(depth, prefix+" "+fmt.Sprintf(format, args...))
	}
	return nil
}

var _ L = (*osSTD)(nil)
