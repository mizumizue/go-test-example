package logger

import "log"

type Logger struct{}

func (l *Logger) Log(v ...interface{}) {
	log.Print(v...)
}

func (l *Logger) Logf(format string, v ...interface{}) {
	log.Print(v...)
}
