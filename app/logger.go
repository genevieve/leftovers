package app

import (
	"fmt"
	"io"
)

type Logger struct {
	newline bool
	writer  io.Writer
}

func NewLogger(writer io.Writer) *Logger {
	return &Logger{
		newline: true,
		writer:  writer,
	}
}

func (l *Logger) clear() {
	if l.newline {
		return
	}

	l.writer.Write([]byte("\n"))
	l.newline = true
}

func (l *Logger) Printf(message string, a ...interface{}) {
	l.clear()
	fmt.Fprintf(l.writer, "%s", fmt.Sprintf(message, a...))
}

func (l *Logger) Prompt(message string) {
	l.clear()
	fmt.Fprintf(l.writer, "%s (y/N): ", message)
	l.newline = true
}
