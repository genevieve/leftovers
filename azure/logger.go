package azure

type logger interface {
	Printf(message string, args ...interface{})
	Println(message string)
	Debugln(message string)
	PromptWithDetails(resourceType, resourceName string) bool
	NoConfirm()
}
