package kms

//go:generate faux --interface logger --output fakes/logger.go
type logger interface {
	Printf(m string, a ...interface{})
	PromptWithDetails(resourceType, resourceName string) (proceed bool)
}
