package writers

import (
	"go.uber.org/zap/zapcore"
)

// Attribute defines a single SGR Code
type Attribute int

// Foreground text colors
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

func getZapLevelColor(level zapcore.Level) Attribute {
	switch level {
	case zapcore.DebugLevel:
		return FgCyan
	case zapcore.InfoLevel:
		return FgBlue
	case zapcore.WarnLevel:
		return FgYellow
	case zapcore.ErrorLevel:
		return FgRed
	case zapcore.DPanicLevel:
		return FgHiRed
	case zapcore.PanicLevel:
		return FgHiBlack
	case zapcore.FatalLevel:
		return FgBlack
	}
	return FgCyan
}
