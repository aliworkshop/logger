package logger

type Logger interface {
	Clone() Logger
	With(Field) Logger
	WithId(id string) Logger
	WithUid(string) Logger
	WithSource(string) Logger
	DebugF(s string, a ...interface{})
	InfoF(s string, a ...interface{})
	WarnF(s string, a ...interface{})
	ErrorF(s string, a ...interface{})
	CriticalF(s string, a ...interface{})
	FatalF(s string, a ...interface{})
}
