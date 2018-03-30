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

	PrintlnCall struct {
		Receives struct {
			Message string
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

func (l *Logger) Println(message string) {
	l.PrintfCall.Receives.Message = message

	l.PrintfCall.Messages = append(l.PrintfCall.Messages, message)
}

func (l *Logger) NoConfirm() {
	l.NoConfirmCall.CallCount++
}
