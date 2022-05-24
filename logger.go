package gokalkan

import "log"

type Logger interface {
	Debug(args ...interface{})
	Error(args ...interface{})
}

type logger struct {
}

func (l *logger) Error(args ...interface{}) {
	log.Println(args...)
}

func (l *logger) Debug(args ...interface{}) {
	log.Println(args...)
}

var _ Logger = (*logger)(nil)

var defaultLogger = &logger{} //nolint:gochecknoglobals
