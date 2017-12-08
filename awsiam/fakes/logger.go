package fakes

import "fmt"

type Logger struct {
	PrintfCall struct {
		Receives struct {
			Message   string
			Arguments []interface{}
		}
		Messages []string
	}

	PromptCall struct {
		Receives struct {
			Message string
		}
	}
}

func (l *Logger) Printf(message string, a ...interface{}) {
	l.PrintfCall.Receives.Message = message
	l.PrintfCall.Receives.Arguments = a

	l.PrintfCall.Messages = append(l.PrintfCall.Messages, fmt.Sprintf(message, a...))
}

func (l *Logger) Prompt(message string) {
	l.PromptCall.Receives.Message = message
}
