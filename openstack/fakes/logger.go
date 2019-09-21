package fakes

type Logger struct {
	PromptWithDetailsCall struct {
		CallCount int
		Receives  struct {
			ResourceType string
			ResourceName string
		}
		ReturnsForCall []LoggerPromptWithDetailsCallReturn
		Returns        LoggerPromptWithDetailsCallReturn
	}
}

type LoggerPromptWithDetailsCallReturn struct {
	Bool bool
}

func (l *Logger) PromptWithDetails(resourceType, resourceName string) bool {
	l.PromptWithDetailsCall.CallCount++
	l.PromptWithDetailsCall.Receives.ResourceType = resourceType
	l.PromptWithDetailsCall.Receives.ResourceName = resourceName

	if len(l.PromptWithDetailsCall.ReturnsForCall) > 0 {
		return l.PromptWithDetailsCall.ReturnsForCall[l.PromptWithDetailsCall.CallCount-1].Bool
	}

	return l.PromptWithDetailsCall.Returns.Bool
}

func (l *Logger) Printf(message string, a ...interface{}) {}
func (l *Logger) Println(message string)                  {}
func (l *Logger) NoConfirm()                              {}
