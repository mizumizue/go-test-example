package logger

import (
	"fmt"
	"testing"
)

// Loggerを追跡する構造体
type SpyLogger struct {
	l   *Logger
	cap string
}

func (sl *SpyLogger) Log(v ...interface{}) {
	sl.l.Log(v...)
	sl.cap = fmt.Sprint(v...)
}

func (sl *SpyLogger) Logf(format string, v ...interface{}) {
	sl.l.Logf(format, v...)
	sl.cap = fmt.Sprintf(format, v...)
}

func SetUpSpyLogger() *SpyLogger {
	return &SpyLogger{
		l: &Logger{},
	}
}

func TestLogger_Log(t *testing.T) {
	sl := SetUpSpyLogger()
	args := []interface{}{"hoge", "foo", "bar"}

	sl.Log(args)

	expected := "[hoge foo bar]"
	if sl.cap != expected {
		t.Fail()
		t.Logf("acctual: %s, expected: %s", sl.cap, expected)
	}
}

func TestLogger_Logf(t *testing.T) {
	sl := SetUpSpyLogger()
	args := []interface{}{"hoge", "foo", "bar"}
	format := "%s - %s - %s"
	sl.Logf(format, args...)

	expected := "hoge - foo - bar"
	if sl.cap != expected {
		t.Fail()
		t.Logf("acctual: %s, expected: %s", sl.cap, expected)
	}
}
