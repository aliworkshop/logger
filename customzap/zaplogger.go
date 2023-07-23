package customzap

import (
	"github.com/aliworkshop/configer"
	"github.com/aliworkshop/loggerlib/logger"
	writer "github.com/aliworkshop/loggerlib/writers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var pcNames map[uintptr]string

func init() {
	pcNames = make(map[uintptr]string)
}

type zapLogger struct {
	loggers []*zap.Logger
	Meta    []logger.Field
	Uid     string
	Source  string
	Id      string
}

func NewLogger(registry configer.Registry, writers []writer.Writer) (logger.Logger, error) {
	config := new(Config)
	err := registry.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	zapLevel := getZapLevel(config.Level)

	if config.Encoding == "" {
		config.Encoding = "json"
	}

	enc := decideEncoder(config.Encoding, zapcore.EncoderConfig{
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
	for _, w := range writers {
		level := zapLevel
		if l, ok := w.Level(); ok {
			level = getZapLevel(l)
		}
		zl.loggers = append(zl.loggers, zap.New(writer.NewCore(enc, zapcore.AddSync(w), zap.NewAtomicLevelAt(level))))
	}
	return zl, err
}

func (zl *zapLogger) Clone() logger.Logger {
	return zl.clone()
}

func (zl *zapLogger) clone() *zapLogger {
	cpy := zl.loggers
	return &zapLogger{loggers: cpy, Meta: zl.Meta, Uid: zl.Uid, Id: zl.Id, Source: zl.Source}
}

func decideEncoder(Type string, config zapcore.EncoderConfig) zapcore.Encoder {
	switch Type {
	case "json":
		return zapcore.NewJSONEncoder(config)
	case "console":
		return zapcore.NewConsoleEncoder(config)
	}
	return nil
}
