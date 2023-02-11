package logger

// Level is the minimum enabled logging level. Note that this is a dynamic
// level, so calling Config.Level.SetLevel will atomically change the log
// level of all loggers descended from this config.
type Level string

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Level = "debug"
	// InfoLevel is the default logging priority.
	InfoLevel Level = "info"
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel Level = "warn"
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel Level = "error"
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel Level = "dpanic"
	// PanicLevel logs a message, then panics.
	PanicLevel Level = "panic"
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel Level = "fatal"
)
