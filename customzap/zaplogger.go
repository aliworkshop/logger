package customzap

import (
	"github.com/aliworkshop/configlib"
	"github.com/aliworkshop/loggerlib/logger"
	"github.com/aliworkshop/loggerlib/writers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var pcNames map[uintptr]string
var pcNamesMtx *sync.RWMutex

func init() {
	pcNames = make(map[uintptr]string)
	pcNamesMtx = new(sync.RWMutex)
}

type metaField struct {
	fields []zapcore.Field
	enc    zapcore.Encoder
}

type zapLogger struct {
	loggers []*zap.Logger
	meta    *metaField
	id      string
	uid     string
	source  string
}

func NewLogger(registry configlib.Registry, writers []writers.Writer) (logger.Logger, error) {
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
	metaEncoder := decideEncoder(config.Encoding, zapcore.EncoderConfig{})
	zl := new(zapLogger)
	zl.meta = &metaField{
		fields: []zapcore.Field{},
		enc:    metaEncoder,
	}
	for _, w := range writers {
		level := zapLevel
		if l, ok := w.Level(); ok {
			level = getZapLevel(l)
		}
		zl.loggers = append(zl.loggers, zap.New(zapcore.NewCore(enc, zapcore.AddSync(w), zap.NewAtomicLevelAt(level))))
	}
	return zl, err
}

func (zl *zapLogger) Clone() logger.Logger {
	return zl.clone()
}

func (zl *zapLogger) clone() *zapLogger {
	zlc := &zapLogger{
		meta:   zl.meta.Clone(),
		id:     zl.id,
		uid:    zl.uid,
		source: zl.source,
	}
	for _, l := range zl.loggers {
		zlc.loggers = append(zlc.loggers, l)
	}
	return zlc
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
