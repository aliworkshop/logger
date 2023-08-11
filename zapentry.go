package logger

import (
	"fmt"
	"go.uber.org/zap/zapcore"
)

func (zl *zapLogger) With(fs Field) Logger {
	return zl.with(fs)
}

func (zl *zapLogger) WithSeverity(severity Severity) *zapLogger {
	zl = zl.with(Field{
		"severity": severity,
	})
	return zl
}

func (zl *zapLogger) WithSource(source string) Logger {
	zl = zl.clone()
	zl.Source = source
	return zl
}

func (zl *zapLogger) WithUid(uid string) Logger {
	zl = zl.clone()
	zl.Uid = uid
	return zl
}

func (zl *zapLogger) WithId(id string) Logger {
	zl = zl.clone()
	zl.Id = id
	return zl
}

func (zl *zapLogger) DebugF(s string, a ...interface{}) {
	zl = zl.WithDefault(SeverityLow)
	for _, l := range zl.loggers {
		l.Debug(fmt.Sprintf(s, a...))
	}
}

func (zl *zapLogger) InfoF(s string, a ...interface{}) {
	zl = zl.WithDefault(SeverityLow)
	for _, l := range zl.loggers {
		l.Info(fmt.Sprintf(s, a...))
	}
}

func (zl *zapLogger) WarnF(s string, a ...interface{}) {
	zl = zl.WithDefault(SeverityMedium)
	for _, l := range zl.loggers {
		l.Warn(fmt.Sprintf(s, a...))
	}
}

func (zl *zapLogger) ErrorF(s string, a ...interface{}) {
	zl = zl.WithDefault(SeverityHigh)
	for _, l := range zl.loggers {
		l.Error(fmt.Sprintf(s, a...))
	}
}

func (zl *zapLogger) CriticalF(s string, a ...interface{}) {
	zl = zl.WithDefault(SeverityCritical)
	for _, l := range zl.loggers {
		l.Error(fmt.Sprintf(s, a...))
	}
}

func (zl *zapLogger) FatalF(s string, a ...interface{}) {
	zl = zl.WithDefault(SeverityHigh)
	for _, l := range zl.loggers {
		l.Fatal(fmt.Sprintf(s, a...))
	}
}

func (zl *zapLogger) WithDefault(severity Severity) *zapLogger {
	zl = zl.with(Field{
		"severity": severity,
		"uuid":     zl.Uid,
		"id":       zl.Id,
		"source":   zl.Source,
	})
	return zl
}

func (zl *zapLogger) with(fs Field) *zapLogger {
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
	zls.Source = zl.Source
	zls.Id = zl.Id
	zls.Uid = zl.Uid
	for _, l := range zl.loggers {
		zls.loggers = append(zls.loggers, l.With(fields...))
	}
	return zls
}
