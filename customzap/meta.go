package customzap

import (
	"aliworkshop/errorslib"
	"github.com/aliworkshop/loggerlib/logger"
	"go.uber.org/zap/zapcore"
)

func (zl *zapLogger) withMeta(fs logger.Field) *zapLogger {
	zl = zl.clone()
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
		case errorslib.ErrorModel:
			fields = append(fields, zapcore.Field{String: v.(errorslib.ErrorModel).Error(), Type: zapcore.StringType, Key: k})
		default:
			fields = append(fields, zapcore.Field{Interface: v, Type: zapcore.ReflectType, Key: k})
		}
	}
	zl.meta.fields = append(zl.meta.fields, fields...)
	return zl
}

func (m *metaField) Clone() *metaField {
	return &metaField{
		enc:    m.enc.Clone(),
		fields: m.fields,
	}
}
