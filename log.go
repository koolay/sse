package sse

import "log"

type DebugLog interface {
	Logf(format string, args ...interface{})
}

type NoLog struct{}

func (l *NoLog) Logf(format string, args ...interface{}) {
}

type StdoutLog struct{}

func (l *StdoutLog) Logf(format string, args ...interface{}) {
	log.Printf(format, args...)
}
