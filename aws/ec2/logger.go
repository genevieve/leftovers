package ec2

//go:generate faux --interface logger --output fakes/logger.go
type logger interface {
	Printf(format string, v ...interface{})
	PromptWithDetails(resourceType, resourceName string) (proceed bool)
}
