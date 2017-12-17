package app

import (
	"fmt"
	"io"
	"strings"
)

type Logger struct {
	newline   bool
	writer    io.Writer
	reader    io.Reader
	noConfirm bool
}

func NewLogger(writer io.Writer, reader io.Reader, noConfirm bool) *Logger {
	return &Logger{
		newline:   true,
		writer:    writer,
		reader:    reader,
		noConfirm: noConfirm,
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
	fmt.Fprintf(l.writer, "\t%s", fmt.Sprintf(message, a...))
}

func (l *Logger) Prompt(message string) bool {
	if l.noConfirm {
		return true
	}

	l.clear()
	fmt.Fprintf(l.writer, "%s (y/N): ", message)
	l.newline = true

	var proceed string
	fmt.Fscanln(l.reader, &proceed)

	proceed = strings.ToLower(proceed)
	if proceed != "yes" && proceed != "y" {
		return false
	}

	return true
}
