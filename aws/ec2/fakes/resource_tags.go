package fakes

type ResourceTags struct {
	DeleteCall struct {
		CallCount int
		Receives  struct {
			FilterName  string
			FilterValue string
		}
		Returns struct {
			Error error
		}
	}
}

func (r *ResourceTags) Delete(filterName, filterValue string) error {
	r.DeleteCall.CallCount++
	r.DeleteCall.Receives.FilterName = filterName
	r.DeleteCall.Receives.FilterValue = filterValue

	return r.DeleteCall.Returns.Error
}
