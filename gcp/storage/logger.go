package storage

//go:generate faux --interface logger --output fakes/logger.go
type logger interface {
	Printf(message string, a ...interface{})
	Debugln(message string)
	PromptWithDetails(resourceType, resourceName string) (proceed bool)
}
