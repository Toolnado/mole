package logger

import (
	"log"
	"os"
)

type Logger struct {
	_error *log.Logger
	_info  *log.Logger
}

func New() *Logger {
	e := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)
	i := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime|log.Lshortfile)
	return &Logger{
		_error: e,
		_info:  i,
	}
}

func (l *Logger) Error(format string, v ...any) {
	l._error.Printf(format, v...)
}

func (l *Logger) Fatal(format string, v ...any) {
	l._error.Fatalf(format, v...)
}

func (l *Logger) Info(format string, v ...any) {
	l._info.Printf(format, v...)
}
