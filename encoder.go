package logger

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

// ShortCallerEncoder serializes a caller in package/file:line format, trimming
// all but the final directory from the full path.
func ShortCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	path := caller.TrimmedPath()
	if atInd := strings.Index(path, "@"); atInd > 0 {
		if slashInd := strings.Index(path[atInd:], "/"); slashInd > atInd {
			path = path[:atInd] + path[atInd+slashInd:]
		}
	}
	enc.AppendString(path)
}
