package noop

import (
	"github.com/aliworkshop/loggerlib/logger"
)

type noop struct{}

func (l *noop) WithId(id string) logger.Logger {
	return l
}

func (l *noop) WithUid(s string) logger.Logger {
	return l
}

func (l *noop) WithSource(s string) logger.Logger {
	return l
}

func (l *noop) Clone() logger.Logger {
	return l
}

func (l *noop) With(field logger.Field) logger.Logger {
	return l
}

func (l *noop) DebugF(s string, a ...interface{}) {
}

func (l *noop) InfoF(s string, a ...interface{}) {
}

func (l *noop) ErrorF(s string, a ...interface{}) {
}

func (l *noop) FatalF(s string, a ...interface{}) {
}

func (l *noop) WarnF(s string, a ...interface{}) {
}

func (l *noop) CriticalF(s string, a ...interface{}) {
}

func NewNoopLogger() logger.Logger {
	return &noop{}
}
