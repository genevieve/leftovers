package artifacts

//go:generate faux --interface logger --output fakes/logger.go
type logger interface {
	Printf(message string, a ...interface{})
	Debugf(message string, a ...interface{})
	PromptWithDetails(resourceType, resourceName string) (proceed bool)
}
