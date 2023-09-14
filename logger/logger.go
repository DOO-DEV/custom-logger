package logger

import (
	"log"
)

type Logger struct {
	logger *log.Logger
}

func New(l *log.Logger) *Logger {
	return &Logger{logger: l}
}

func (l Logger) Warn(msg string) {
	l.logger.SetPrefix("1: ")
	l.logger.Println(msg)
}

func (l Logger) Info(msg string) {
	l.logger.SetPrefix("2: ")
	l.logger.Println(msg)
}

func (l Logger) Error(msg string) {
	l.logger.SetPrefix("3: ")
	l.logger.Println(msg)
}
