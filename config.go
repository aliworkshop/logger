package logger

import (
	"github.com/aliworkshop/logger/writers"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	// Level is the minimum enabled logging level. Note that this is a dynamic
	// level, so calling Config.Level.SetLevel will atomically change the log
	// level of all loggers descended from this config.
	Level writers.Level
	// Development puts the logger in development mode, which changes the
	// behavior of DPanicLevel and takes stacktraces more liberally.
	Development bool
	// Encoding sets the logger's encoding. Valid values are "json" and
	// "console", as well as any third-party encodings registered via
	// RegisterEncoder.
	Encoding string
}

func getZapLevel(level writers.Level) zapcore.Level {
	switch level {
	case writers.DebugLevel:
		return zapcore.DebugLevel
	case writers.InfoLevel:
		return zapcore.InfoLevel
	case writers.WarnLevel:
		return zapcore.WarnLevel
	case writers.ErrorLevel:
		return zapcore.ErrorLevel
	case writers.DPanicLevel:
		return zapcore.DPanicLevel
	case writers.PanicLevel:
		return zapcore.PanicLevel
	case writers.FatalLevel:
		return zapcore.FatalLevel
	}
	return zapcore.DebugLevel
}
