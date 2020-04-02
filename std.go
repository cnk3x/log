package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

func New() L {
	return &osSTD{
		l:     log.New(os.Stderr, "", 0),
		level: DebugLevel,
	}
}

type osSTD struct {
	l     *log.Logger
	name  string
	level Level
}

func (s *osSTD) SetLevel(level string) {
	s.level = ParseLevel(level)
}

func (s *osSTD) SetName(name string) {
	s.name = name
}

func (s *osSTD) Output(depth int, level Level, format string, args ...interface{}) error {
	if level >= s.level {
		prefix := fmt.Sprintf("%s<%s> %s %s",
			s.name,
			level.String()[:1],
			time.Now().Format(time.RFC3339),
			format,
		)
		return s.l.Output(depth, fmt.Sprintf(prefix, args...))
	}
	return nil
}

var _ L = (*osSTD)(nil)
