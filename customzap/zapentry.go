package customzap

import (
	"fmt"
	"github.com/aliworkshop/loggerlib/logger"
	"go.uber.org/zap/zapcore"
	"runtime"
	"strings"
)

func (zl *zapLogger) With(fs logger.Field) logger.Logger {
	return zl.withMeta(fs)
}

func (zl *zapLogger) WithSeverity(severity logger.Severity) *zapLogger {
	zl = zl.with(logger.Field{
		"severity": severity,
	})
	return zl
}

func (zl *zapLogger) WithSource(source string) logger.Logger {
	zl = zl.clone()
	zl.Source = source
	return zl
}

func (zl *zapLogger) WithUid(uid string) logger.Logger {
	zl = zl.clone()
	zl.Uid = uid
	return zl
}

func (zl *zapLogger) WithId(id string) logger.Logger {
	zl = zl.clone()
	zl.Id = id
	return zl
}

func (zl *zapLogger) DebugF(s string, a ...interface{}) {
	zl = zl.WithDefault(logger.SeverityLow)
	for _, l := range zl.loggers {
		l.Debug(fmt.Sprintf(s, a...))
	}
}

func (zl *zapLogger) InfoF(s string, a ...interface{}) {
	zl = zl.WithDefault(logger.SeverityLow)
	for _, l := range zl.loggers {
		l.Info(fmt.Sprintf(s, a...))
	}
}

func (zl *zapLogger) WarnF(s string, a ...interface{}) {
	zl = zl.WithDefault(logger.SeverityMedium)
	for _, l := range zl.loggers {
		l.Warn(fmt.Sprintf(s, a...))
	}
}

func (zl *zapLogger) ErrorF(s string, a ...interface{}) {
	zl = zl.WithDefault(logger.SeverityHigh)
	for _, l := range zl.loggers {
		l.Error(fmt.Sprintf(s, a...))
	}
}

func (zl *zapLogger) CriticalF(s string, a ...interface{}) {
	zl = zl.WithDefault(logger.SeverityCritical)
	for _, l := range zl.loggers {
		l.Error(fmt.Sprintf(s, a...))
	}
}

func (zl *zapLogger) FatalF(s string, a ...interface{}) {
	zl = zl.WithDefault(logger.SeverityHigh)
	for _, l := range zl.loggers {
		l.Fatal(fmt.Sprintf(s, a...))
	}
}

func (zl *zapLogger) WithDefault(severity logger.Severity) *zapLogger {
	zl = zl.with(logger.Field{
		"severity": severity,
		"context":  getContext(),
		"meta":     zl.Meta,
		"UID":      zl.Uid,
		"Id":       zl.Id,
		"source":   zl.Source,
	})
	return zl
}

func (zl *zapLogger) with(fs logger.Field) *zapLogger {
	if len(fs) == 0 {
		return zl
	}
	fields := make([]zapcore.Field, 0)
	for k, v := range fs {
		switch v.(type) {
		case string:
			fields = append(fields, zapcore.Field{String: v.(string), Type: zapcore.StringType, Key: k})
		case int:
			fields = append(fields, zapcore.Field{Integer: int64(v.(int)), Type: zapcore.Int64Type, Key: k})
		case error:
			fields = append(fields, zapcore.Field{Interface: v.(error), Type: zapcore.ErrorType, Key: k})
		case bool:
			fields = append(fields, zapcore.Field{Interface: v.(bool), Type: zapcore.BoolType, Key: k})
		default:
			fields = append(fields, zapcore.Field{Interface: v, Type: zapcore.ReflectType, Key: k})
		}
	}
	zls := new(zapLogger)
	for _, l := range zl.loggers {
		zls.loggers = append(zls.loggers, l.With(fields...))
	}
	return zls
}

func getContext() string {
	ptr := make([]uintptr, 3)
	runtime.Callers(4, ptr)
	context := ""
	for i := len(ptr) - 1; i >= 0; i-- {
		p := ptr[i]
		if name, ok := pcNames[p]; ok {
			context += name + "|"
		} else {
			f := runtime.FuncForPC(p)
			if f != nil {
				pcNames[p] = handleName(f.Name())
				context += pcNames[p] + "|"
			}
		}
	}
	return strings.TrimSuffix(context, "|")
}

func handleName(name string) string {
	split := strings.Split(name, "/")
	split = strings.Split(split[len(split)-1], ".")
	return split[len(split)-1]
}
