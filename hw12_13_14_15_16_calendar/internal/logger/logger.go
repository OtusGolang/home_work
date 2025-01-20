package logger

import "fmt"

type Logger struct { // TODO
}

func New(level string) *Logger {
	return &Logger{}
}

func (l Logger) Info(msg string) {
	fmt.Println(msg)
}

func (l Logger) Error(msg string) {
	// TODO
}

// TODO
