package customzap

import (
	"github.com/aliworkshop/logger/logger"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	// Level is the minimum enabled logging level. Note that this is a dynamic
	// level, so calling Config.Level.SetLevel will atomically change the log
	// level of all loggers descended from this config.
	Level logger.Level
	// Development puts the logger in development mode, which changes the
	// behavior of DPanicLevel and takes stacktraces more liberally.
	Development bool
	// Encoding sets the logger's encoding. Valid values are "json" and
	// "console", as well as any third-party encodings registered via
	// RegisterEncoder.
	Encoding string
}

func getZapLevel(level logger.Level) zapcore.Level {
	switch level {
	case logger.DebugLevel:
		return zapcore.DebugLevel
	case logger.InfoLevel:
		return zapcore.InfoLevel
	case logger.WarnLevel:
		return zapcore.WarnLevel
	case logger.ErrorLevel:
		return zapcore.ErrorLevel
	case logger.DPanicLevel:
		return zapcore.DPanicLevel
	case logger.PanicLevel:
		return zapcore.PanicLevel
	case logger.FatalLevel:
		return zapcore.FatalLevel
	}
	return zapcore.DebugLevel
}
