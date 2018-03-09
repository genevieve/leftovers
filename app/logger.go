package app

import (
	"fmt"
	"io"
	"strings"
	"sync"
)

type Logger struct {
	newline   bool
	writer    io.Writer
	mutex     *sync.Mutex
	reader    io.Reader
	noConfirm bool
}

func NewLogger(writer io.Writer, reader io.Reader, noConfirm bool) *Logger {
	return &Logger{
		newline:   true,
		writer:    writer,
		mutex:     &sync.Mutex{},
		reader:    reader,
		noConfirm: noConfirm,
	}
}

// clear is not threadsafe.
func (l *Logger) clear() {
	if l.newline {
		return
	}

	l.writer.Write([]byte("\n"))
	l.newline = true
}

func (l *Logger) Printf(message string, a ...interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.clear()
	fmt.Fprintf(l.writer, "\t%s", fmt.Sprintf(message, a...))
}

func (l *Logger) Println(message string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.clear()
	fmt.Fprintf(l.writer, "%s\n\n", message)
}

// Please note that Prompt will block all other goroutines
// attempting to print to this logger while
// waiting for user input.
func (l *Logger) Prompt(message string) bool {
	if l.noConfirm {
		return true
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()
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

func (l *Logger) NoConfirm() {
	l.noConfirm = true
}
