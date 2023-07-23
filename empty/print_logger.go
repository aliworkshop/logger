package empty

import (
	"fmt"
	"github.com/aliworkshop/logger/logger"
)

type emptyLogger struct{}

func (l *emptyLogger) WithId(id string) logger.Logger {
	return l
}

func (l *emptyLogger) WithUid(s string) logger.Logger {
	return l
}

func (l *emptyLogger) WithSource(s string) logger.Logger {
	return l
}

func (l *emptyLogger) Clone() logger.Logger {
	return l
}

func (l *emptyLogger) With(field logger.Field) logger.Logger {
	return l
}

func (l *emptyLogger) DebugF(s string, a ...interface{}) {
	fmt.Printf(s, a...)
}

func (l *emptyLogger) InfoF(s string, a ...interface{}) {
	fmt.Printf(s, a...)
}

func (l *emptyLogger) ErrorF(s string, a ...interface{}) {
	fmt.Printf(s, a...)
}

func (l *emptyLogger) FatalF(s string, a ...interface{}) {
	fmt.Printf(s, a...)
}

func (l *emptyLogger) WarnF(s string, a ...interface{}) {
	fmt.Printf(s, a...)
}

func (l *emptyLogger) CriticalF(s string, a ...interface{}) {
	fmt.Printf(s, a...)
}

func NewEmptyLogger() logger.Logger {
	return &emptyLogger{}
}
