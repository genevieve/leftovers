package sql

type logger interface {
	Printf(message string, a ...interface{})
	Debugln(message string)
	PromptWithDetails(resourceType, resourceName string) bool
}
