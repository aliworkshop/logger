package customzap

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/aliworkshop/loggerlib/logger"
	"go.uber.org/zap/zapcore"
)

func (zl *zapLogger) With(fs logger.Field) logger.Logger {
	return zl.withMeta(fs)
}

func (zl *zapLogger) WithSource(source string) logger.Logger {
	zl = zl.clone()
	zl.source = source
	return zl
}

func (zl *zapLogger) WithUid(uid string) logger.Logger {
	zl = zl.clone()
	zl.uid = uid
	return zl
}

func (zl *zapLogger) WithId(id string) logger.Logger {
	zl = zl.clone()
	zl.id = id
	return zl
}

func (zl *zapLogger) withDefault(severity logger.Severity) *zapLogger {
	zl = zl.clone()
	data, err := zl.meta.enc.EncodeEntry(zapcore.Entry{}, zl.meta.fields)
	if err != nil {
		fmt.Println(err)
		return zl
	}
	for i, l := range zl.loggers {
		zl.loggers[i] = l.With()
	}
	zl.addToCore(
		zapcore.Field{String: string(severity), Type: zapcore.StringType, Key: "severity"},
		zapcore.Field{String: getContext(), Type: zapcore.StringType, Key: "context"},
		zapcore.Field{String: zl.id, Type: zapcore.StringType, Key: "ID"},
		zapcore.Field{String: zl.source, Type: zapcore.StringType, Key: "Source"},
		zapcore.Field{String: zl.uid, Type: zapcore.StringType, Key: "UID"},
		zapcore.Field{Key: "meta", Type: zapcore.ByteStringType, Interface: data.Bytes()},
	)
	return zl
}

func (zl *zapLogger) addToCore(fields ...zapcore.Field) {
	for i, l := range zl.loggers {
		zl.loggers[i] = l.With(fields...)
	}
}

func getContext() string {
	ptr := make([]uintptr, 3)
	runtime.Callers(4, ptr)
	context := ""
	for i := len(ptr) - 1; i >= 0; i-- {
		p := ptr[i]
		pcNamesMtx.RLock()
		name, ok := pcNames[p]
		pcNamesMtx.RUnlock()
		if ok {
			context += name + "|"
		} else {
			f := runtime.FuncForPC(p)
			if f != nil {
				pcNamesMtx.Lock()
				pcNames[p] = handleName(f.Name())
				pcNamesMtx.Unlock()
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

func (zl *zapLogger) DebugF(s string, a ...interface{}) {
	zl = zl.withDefault(logger.SeverityLow)
	for _, l := range zl.loggers {
		l.Debug(fmt.Sprintf(s, a...))
	}
}

func (zl *zapLogger) InfoF(s string, a ...interface{}) {
	zl = zl.withDefault(logger.SeverityLow)
	for _, l := range zl.loggers {
		l.Info(fmt.Sprintf(s, a...))
	}
}

func (zl *zapLogger) WarnF(s string, a ...interface{}) {
	zl = zl.withDefault(logger.SeverityMedium)
	for _, l := range zl.loggers {
		l.Warn(fmt.Sprintf(s, a...))
	}
}

func (zl *zapLogger) ErrorF(s string, a ...interface{}) {
	zl = zl.withDefault(logger.SeverityHigh)
	for _, l := range zl.loggers {
		l.Error(fmt.Sprintf(s, a...))
	}
}

func (zl *zapLogger) CriticalF(s string, a ...interface{}) {
	zl = zl.withDefault(logger.SeverityCritical)
	for _, l := range zl.loggers {
		l.Error(fmt.Sprintf(s, a...))
	}
}

func (zl *zapLogger) FatalF(s string, a ...interface{}) {
	zl = zl.withDefault(logger.SeverityHigh)
	for _, l := range zl.loggers {
		l.Fatal(fmt.Sprintf(s, a...))
	}
}
