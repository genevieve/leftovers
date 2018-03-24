package fakes

type Logger struct {
	PromptWithDetailsCall struct {
		CallCount int
		Receives  struct {
			Type string
			Name string
		}
		Returns struct {
			Proceed bool
		}
	}
}

func (l *Logger) PromptWithDetails(resourceType, resourceName string) bool {
	l.PromptWithDetailsCall.CallCount++
	l.PromptWithDetailsCall.Receives.Type = resourceType
	l.PromptWithDetailsCall.Receives.Name = resourceName

	return l.PromptWithDetailsCall.Returns.Proceed
}
