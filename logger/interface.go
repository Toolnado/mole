package logger

type Logger interface {
	Error(v ...any)
	Fatal(v ...any)
	Info(v ...any)
}
