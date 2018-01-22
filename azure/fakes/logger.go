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
		CallCount int
		Receives  struct {
			Message string
		}
		Returns struct {
			Proceed bool
		}
	}
}

func (l *Logger) Printf(message string, a ...interface{}) {
	l.PrintfCall.Receives.Message = message
	l.PrintfCall.Receives.Arguments = a

	l.PrintfCall.Messages = append(l.PrintfCall.Messages, fmt.Sprintf(message, a...))
}

func (l *Logger) Prompt(message string) bool {
	l.PromptCall.CallCount++
	l.PromptCall.Receives.Message = message

	return l.PromptCall.Returns.Proceed
}
