package kms

type logger interface {
	PromptWithDetails(resourceType, resourceName string) bool
}
