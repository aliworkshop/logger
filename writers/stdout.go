package writers

import (
	"github.com/aliworkshop/logger/logger"
	"io"
	"os"
)

type stdout struct {
	level string
	io.Writer
}

func newStdout(level string) Writer {
	return &stdout{
		level:  level,
		Writer: os.Stdout,
	}
}

func (s *stdout) Level() (logger.Level, bool) {
	if s.level == "" {
		return "", false
	}
	return logger.Level(s.level), true
}

type stdErr struct {
	level string
	io.Writer
}

func newStderr(level string) Writer {
	return &stdErr{
		level:  level,
		Writer: os.Stderr,
	}
}

func (s *stdErr) Level() (logger.Level, bool) {
	if s.level == "" {
		return "", false
	}
	return logger.Level(s.level), true
}
