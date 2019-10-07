package azure

//go:generate faux --interface logger --output fakes/logger.go
type logger interface {
	Printf(message string, args ...interface{})
	Println(message string)
	Debugln(message string)
	PromptWithDetails(resourceType, resourceName string) (proceed bool)
	NoConfirm()
}
