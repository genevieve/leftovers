package container

type logger interface {
	Printf(message string, a ...interface{})
	Debugf(message string, a ...interface{})
	PromptWithDetails(resourceType, resourceName string) bool
}
