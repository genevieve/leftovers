package fakes

type ResourceTags struct {
	DeleteCall struct {
		CallCount int
		Receives  struct {
			ResourceType string
			ResourceId   string
		}
		Returns struct {
			Error error
		}
	}
}

func (r *ResourceTags) Delete(resourceType, resourceId string) error {
	r.DeleteCall.CallCount++
	r.DeleteCall.Receives.ResourceType = resourceType
	r.DeleteCall.Receives.ResourceId = resourceId

	return r.DeleteCall.Returns.Error
}
