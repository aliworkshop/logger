package logger

import (
	"github.com/aliworkshop/configer"
	"github.com/aliworkshop/logger/writers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type zapLogger struct {
	loggers []*zap.Logger
	Uid     string
	Source  string
	Id      string
}

func NewSimpleLogger(level writers.Level, encoding string) Logger {
	zapLevel := getZapLevel(level)

	enc := decideEncoder(encoding, zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   ShortCallerEncoder,
	})
	zl := new(zapLogger)
	zl.loggers = append(zl.loggers, zap.New(writers.NewCore(enc, zapcore.AddSync(os.Stdout), zap.NewAtomicLevelAt(zapLevel))))
	return zl
}

func NewLogger(registry configer.Registry, wrts []writers.Writer) (Logger, error) {
	cfg := new(Config)
	err := registry.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	zapLevel := getZapLevel(cfg.Level)

	if cfg.Encoding == "" {
		cfg.Encoding = "json"
	}

	enc := decideEncoder(cfg.Encoding, zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   ShortCallerEncoder,
	})
	zl := new(zapLogger)
	for _, w := range wrts {
		level := zapLevel
		if l, ok := w.Level(); ok {
			level = getZapLevel(l)
		}
		zl.loggers = append(zl.loggers, zap.New(writers.NewCore(enc, zapcore.AddSync(w), zap.NewAtomicLevelAt(level))))
	}
	return zl, err
}

func (zl *zapLogger) Clone() Logger {
	return zl.clone()
}

func (zl *zapLogger) clone() *zapLogger {
	cpy := zl.loggers
	return &zapLogger{loggers: cpy, Uid: zl.Uid, Id: zl.Id, Source: zl.Source}
}

func decideEncoder(Type string, config zapcore.EncoderConfig) zapcore.Encoder {
	switch Type {
	case JsonEncoding:
		return zapcore.NewJSONEncoder(config)
	case ConsoleEncoding:
		return zapcore.NewConsoleEncoder(config)
	}
	return nil
}
