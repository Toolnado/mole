package logger

import (
	"log"
	"os"
)

type DefaultLogger struct {
	_error *log.Logger
	_info  *log.Logger
}

func New() *DefaultLogger {
	e := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)
	i := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime|log.Lshortfile)

	return &DefaultLogger{
		_error: e,
		_info:  i,
	}
}

func (l *DefaultLogger) Error(v ...any) {
	l._error.Printf(v[0].(string), v[1:])
}

func (l *DefaultLogger) Fatal(v ...any) {
	l._error.Fatalf(v[0].(string), v[1:])
}

func (l *DefaultLogger) Info(v ...any) {
	l._info.Printf(v[0].(string), v[1:])
}
