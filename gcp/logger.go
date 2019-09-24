package gcp

type logger interface {
	Printf(message string, a ...interface{})
	Println(message string)
	Debugf(message string, a ...interface{})
	Debugln(message string)
	PromptWithDetails(resourceType, resourceName string) bool
	NoConfirm()
}
