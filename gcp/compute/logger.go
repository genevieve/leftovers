package compute

//go:generate faux --interface logger --output fakes/logger.go
type logger interface {
	Printf(m string, a ...interface{})
	Debugf(message string, a ...interface{})
	Debugln(message string)
	PromptWithDetails(resourceType, resourceName string) (proceed bool)
}
