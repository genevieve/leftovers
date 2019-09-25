package compute

type logger interface {
	Printf(m string, a ...interface{})
	Debugf(message string, a ...interface{})
	Debugln(message string)
	PromptWithDetails(resourceType, resourceName string) bool
}
