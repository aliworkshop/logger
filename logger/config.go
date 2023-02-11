package logger

type Config struct {
	// Type is logger type, Valid values are "zap"
	Type string
	Writers map[string]interface{}
}
